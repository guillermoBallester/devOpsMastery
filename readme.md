# DevOps Mastery

A project demonstrating mastery of DevOps skills, Golang back-end development, Kubernetes, and cloud services.

## Features

- Golang HTTP server with Chi router
- Connection management for tracking and graceful shutdown
- Configurable server settings
- Docker and Kubernetes deployment ready
- CI/CD with GitHub Actions

## Getting Started

### Prerequisites

- Go 1.23 or later
- Make
- Docker (for containerized deployment)

### Installation

#### Running Locally

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

The server will start on port 8081 (or the port specified in your config.yaml).

#### Using Docker

1. Build and run with Docker Compose:
   ```sh
   docker-compose up -d
   ```

2. Or build and run manually:
   ```sh
   docker build -t devops-mastery:latest .
   docker run -p 8081:8081 devops-mastery:latest
   ```

## API Endpoints

- **Health Check**: `GET /health`
- **Hello World**: `GET /api/v1/hello`

## Configuration

Configuration is handled through `config.yaml`. You can override settings with environment variables using the `APP_` prefix.

```yaml
server:
  http:
    port: 8081
    readTimeout: 15s
    writeTimeout: 15s
    idleTimeout: 60s
    shutdownTimeout: 30s
```

Example with environment variables:
```sh
# Set HTTP port via environment variable
export APP_SERVER_HTTP_PORT=9090
make run
```

## Development

### Available Commands

#### Go Development
- `make build` - Build the application
- `make run` - Run the application
- `make test` - Run tests
- `make clean` - Clean build artifacts
- `make deps` - Install dependencies
- `make fmt` - Format the code
- `make vet` - Vet the code

#### Docker Commands
- `make docker-build` - Build the Docker image
- `make docker-run` - Build and run in a Docker container
- `make docker-stop` - Stop and remove the container

#### Docker Compose Commands
- `make docker-compose-up` - Start services with Docker Compose
- `make docker-compose-down` - Stop Docker Compose services
- `make docker-compose-logs` - View logs from running services

## Architecture

### Connection Management

The application includes a connection management system that:
- Tracks all active HTTP connections

### Server Lifecycle

1. Server initialization with configuration
2. Connection tracking through middleware
3. Graceful shutdown on termination signals (SIGTERM/SIGINT)

## Future Enhancements

- Implement OpenTelemetry for observability
- Add Prometheus metrics for monitoring
- Expand test coverage