# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the Go app with CGO disabled (no need for SQLite C bindings)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# Final stage - use distroless for minimal attack surface
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Run the app (distroless already uses non-root user)
CMD ["./main"]
