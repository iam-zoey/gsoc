# Use the official Golang image as the base image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod ./

# Download the Go modules dependencies
RUN go mod download

# Copy the rest of the application code
COPY src .

# Build the WebAssembly binary and the HTTP server
RUN GOOS=js GOARCH=wasm go build -o main.wasm main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the HTTP server
CMD ["go", "run", "server.go"]
