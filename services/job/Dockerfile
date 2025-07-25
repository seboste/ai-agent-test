# Build stage
FROM golang:1.21-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy source code including vendor directory
COPY . .

# Build the binary using vendor directory
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o job main.go

# Runtime stage - using scratch for minimal image
FROM scratch

# Copy ca-certificates from the builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary from builder stage
COPY --from=builder /app/job /job

# Set default PORT environment variable
ENV PORT=8080

# Expose the port
EXPOSE $PORT

# Run the binary
CMD ["/job"]