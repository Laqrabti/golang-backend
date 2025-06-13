# Golang Backend API

[![CI](https://github.com/Laqrabti/golang-backend/actions/workflows/go-ci.yml/badge.svg)](https://github.com/Laqrabti/golang-backend/actions/workflows/go-ci.yml)
[![Code Coverage](https://codecov.io/gh/Laqrabti/golang-backend/branch/main/graph/badge.svg)](https://codecov.io/gh/Laqrabti/golang-backend)
[![Docker Build](https://img.shields.io/docker/cloud/build/laqrabti/golang-backend)](https://hub.docker.com/r/laqrabti/golang-backend)
[![Go Report Card](https://goreportcard.com/badge/github.com/Laqrabti/golang-backend)](https://goreportcard.com/report/github.com/Laqrabti/golang-backend)

A production-ready Go backend with Gin framework featuring:

- REST API endpoints
- CORS support
- Health checks
- Docker containerization
- Unit tests with 85%+ coverage
- CI/CD pipeline
- Security scanning
- Linting

## API Endpoints

-  - Service health check
-  - Simple ping/pong
-  - API status
-  - Echo service with message reversal

## Running Locally

```bash
go run main.go
```

## Building

```bash
go build -o backend
```

## Testing

```bash
go test -v ./...
```

## Docker Build

```bash
docker build -t golang-backend .
docker run -p 8080:8080 golang-backend
```
