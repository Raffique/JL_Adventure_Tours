# Use a Go version compatible with Air
FROM golang:1.23 AS base

WORKDIR /app

# Install Air from the new module path
RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Expose application port
EXPOSE 8080

# Use Air for live reloading during development
CMD ["air", "-c", ".air.toml"]
