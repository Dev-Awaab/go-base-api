# Use the official Golang image for building the application
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o server cmd/server/main.go

# Use a minimal base image for the final container
FROM debian:bullseye-slim

# Set environment variables
ENV PORT=8080
ENV SERVER_ADDRESS=0.0.0.0:$PORT

# Install dependencies (e.g., CA certificates for HTTPS)
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/server .

# Copy the .env file (if needed for the app)
COPY .env .env

# Expose the application port
EXPOSE 8080

# Run the Go application
CMD ["./server"]
