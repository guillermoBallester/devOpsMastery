package telemetry

import (
	"context"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Telemetry struct {
	Logger         *otelzap.Logger
	TracerProvider *sdktrace.TracerProvider
}

// Initialize initializes both tracing and logging with OpenTelemetry
func Initialize(ctx context.Context, serviceName, serviceVersion, endpoint string, isDev bool) (*Telemetry, error) {
	tracer, err := setUpTracer(ctx, serviceName, serviceVersion, endpoint)
	if err != nil {
		return nil, err
	}

	logger, err := setUpLogger(isDev)
	if err != nil {
		return nil, err
	}

	return &Telemetry{
		Logger:         logger,
		TracerProvider: tracer,
	}, nil

}

// Shutdown cleans up telemetry resources
func (t *Telemetry) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := t.Logger.Sync(); err != nil {
		return err
	}

	if err := t.TracerProvider.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func setUpTracer(ctx context.Context, serviceName, serviceVersion, endpoint string) (*sdktrace.TracerProvider, error) {
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		),
	)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return tracerProvider, nil
}

func setUpLogger(isDev bool) (*otelzap.Logger, error) {
	var zapLogger *zap.Logger
	var err error
	if isDev {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapLogger, err = config.Build()
	} else {
		config := zap.NewProductionConfig()
		config.OutputPaths = []string{"stdout"}
		zapLogger, err = config.Build()
	}

	if err != nil {
		return nil, err
	}

	logger := otelzap.New(zapLogger,
		otelzap.WithStackTrace(true),
		otelzap.WithMinLevel(zapcore.InfoLevel),
	)

	return logger, nil
}
