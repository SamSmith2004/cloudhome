# ---- Build Stage ----
FROM golang:1.23.7-alpine AS builder

# Environment variables
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /app

# Copy go module files and download dependencies first to leverage caching
COPY api/go.mod api/go.sum ./
RUN go mod download

# Copy the api source code
COPY api/ ./

# Build + output to /app/server inside the build stage
RUN go build -ldflags="-w -s" -o /app/server ./cmd/server

# ---- Runtime Stage ----
FROM alpine:latest
WORKDIR /app

# Copy binary
COPY --from=builder /app/server /app/server

# Expose port 8080
EXPOSE 8080

# Run
CMD ["/app/server"]
