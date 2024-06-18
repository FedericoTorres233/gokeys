# Start with a Golang base image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod .
COPY go.sum .

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN make build

# Command to run the executable
CMD ["./bin/gokeys"]
