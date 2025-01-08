# Use a lightweight Go image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod file to container
COPY go.mod ./

# Download dependencies (skip go.sum)
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application
RUN go build -o paxos-service .

# Expose port 8080 for the service
EXPOSE 8080

# Run the application
CMD ["./paxos-service"]
