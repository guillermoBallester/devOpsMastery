package service

import (
	"context"
	"github.com/guillermoBallester/devOpsMastery/src/internal/telemetry"
)

type HelloService struct {
	tl *telemetry.Telemetry
}

func NewHelloService(tl *telemetry.Telemetry) *HelloService {
	return &HelloService{
		tl: tl,
	}
}

func (s *HelloService) GetHello(ctx context.Context) (string, error) {
	if ctx.Err() != nil {
		s.tl.Logger.Error(ctx.Err().Error())
		return "", ctx.Err()
	}
	return "Hello World!", nil
}
