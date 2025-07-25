# Job Service

A microservice for managing compute jobs in the Green Load Shifting Platform.

## Features

- REST API for job management
- Configurable port via `PORT` environment variable
- Health check endpoint
- Graceful shutdown support
- Minimal Docker image (< 7MB)

## Building

### Local Build
```bash
go build -o job main.go
```

### Docker Build
```bash
docker build -t job-service .
```

## Running

### Local Run
```bash
# Default port 8080
./job

# Custom port
PORT=9000 ./job
```

### Docker Run
```bash
# Default port 8080
docker run -p 8080:8080 job-service

# Custom port
docker run -p 9000:9000 -e PORT=9000 job-service
```

## API Endpoints

- `GET /health` - Health check endpoint
- `GET /jobs` - List jobs (placeholder implementation)
- `POST /jobs` - Create job (placeholder implementation)

## Environment Variables

- `PORT` - Server port (default: 8080)

## Docker Image Details

- Base image: `scratch` (minimal footprint)
- Size: ~7MB
- Includes CA certificates for HTTPS support
- Runs as non-root user in production builds