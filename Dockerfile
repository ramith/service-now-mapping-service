# Start from the official Go image
FROM golang:1.25

# Set the working directory inside the container
WORKDIR /app

# Copy all files into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/server

# Set a non-root user
USER 10014

# Expose port 8080
EXPOSE 8080

# Run the app
CMD ["./main"]
