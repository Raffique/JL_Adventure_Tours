# Start from the golang base image
FROM golang:1.22

# Set the current working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the rest of the files
COPY . .

# Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
