# Stage 1: Build
FROM golang:1.26-alpine AS builder

# Set working directory
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
# We target cmd/api/main.go as the entry point
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Stage 2: Final image
FROM alpine:latest

# Install certificates for HTTPS requests if needed
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy .env file if it exists (Optional, environment variables are preferred in Docker)
# COPY .env .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
