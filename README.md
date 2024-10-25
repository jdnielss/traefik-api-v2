# My Go API

This is a simple Go API that serves a "Hello, World!" message and a health check endpoint.

## Endpoints

- `GET /`: Returns "Hello, World!"
- `GET /health`: Returns "OK"

## Running the API

### Using Docker

1. Build the Docker image:
   ```sh
   docker build -t my-go-api .
