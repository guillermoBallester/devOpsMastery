# DevOps Mastery

A project demonstrating mastery of DevOps skills, Golang back-end development, Kubernetes, and cloud services.

## Getting Started

### Prerequisites

- Go 1.24 or later
- Make

### Installation

1. Clone the repository
   ```sh
   git clone https://github.com/guillermoBallester/devOpsMastery.git
   cd devOpsMastery
   ```

2. Install dependencies
   ```sh
   make deps
   ```

3. Build the project
   ```sh
   make build
   ```

4. Run the server
   ```sh
   make run
   ```

The server will start on port 8080 (or the port specified in your config.yaml).

## API Endpoints

- **Health Check**: `GET /health`
- **Hello World**: `GET /api/v1/hello`

## Development

### Available Commands

- `make build` - Build the application
- `make run` - Run the application
- `make test` - Run tests
- `make clean` - Clean build artifacts
- `make deps` - Install dependencies
- `make fmt` - Format the code
- `make vet` - Vet the code

## Configuration

Configuration is handled through `config.yaml`. You can override settings with environment variables using the `APP_` prefix.

Example:
```sh
# Set HTTP port via environment variable
export APP_SERVER_HTTP_PORT=9090
make run
```

## Future Enhancements

- Add gRPC API
- Add database integration with PostgreSQL
- Implement OpenTelemetry for observability
- Add Kubernetes deployment manifests
- Set up CI/CD with GitHub Actions
- Integrate with GCP and GKE