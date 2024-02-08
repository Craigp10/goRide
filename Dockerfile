# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o route-raydar

# Expose the ports for gRPC and HTTP servers
EXPOSE 50051 8080

# Command to run the application
CMD ["./route-raydar"]
