# command to run the containerized backend: docker run -p 8080:8080 backend (run in the backend dir.)

# Build stage
FROM golang:1.24.5 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Final stage: minimal image
FROM debian:bookworm-slim

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your app runs on (change if needed)
EXPOSE 8080

# Run the binary
CMD ["./main"]