# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./main"]
