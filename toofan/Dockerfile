# Use the official Golang image from the Docker Hub
FROM golang:1.21-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code from the app folder
COPY app/. .

# Build the Go application
RUN go build -o main .

# Expose port 8003
EXPOSE 8003

# Command to run the application
CMD ["./main"]
