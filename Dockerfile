# Stage 1: Build the Go application
FROM golang:1.24.2-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download all Go modules. This helps in caching dependencies.
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application.
# Disable CGO for static compilation for smaller images and fewer dependencies.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Create a minimal image
FROM alpine:3.22.0

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your Gin application listens on (default is 8080)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
