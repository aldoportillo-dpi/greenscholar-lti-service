# Go Backend Template

This repository provides a template for building backend services in Go. It includes a structured setup with Docker support, API routing using Chi, and OpenAPI integration.

## Features

- Go 1.24 based application
- Docker multi-stage build for smaller production images
- Chi router for HTTP API handling
- OpenAPI/Swagger integration for API documentation
- Database migrations support

## Getting Started

### Prerequisites

- Go 1.24 or later
- Docker and Docker Compose (optional)
- Git

### Development

1. Clone this repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

### Building With Docker

```bash
# Build the Docker image
docker build -t go-backend-service --build-arg GITHUB_TOKEN=your_github_token .

# Run the container
docker run -p 80:80 go-backend-service
```

## Project Structure

- `/common` - Shared utilities and common functionality
- `/db/migrations` - Database migration files
- `/openapi` - OpenAPI/Swagger specifications and generated code

## API Documentation

API documentation is generated from the OpenAPI specification in `/openapi/openapi.yaml`. The API structure is accessible when the service is running.

## Database

The service includes support for database migrations that are copied into the Docker image during build.