# Use official Golang image as the base
FROM golang:1.21 AS builder

# Set working directory inside container
WORKDIR /app

# Copy Go modules and dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o chat-system

# Use a minimal base image for production
FROM alpine:latest  

# Set working directory inside the minimal image
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/chat-system .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./chat-system"]

