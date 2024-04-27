# Use an official Go runtime as a parent image for backend
FROM golang:latest AS backend

# Set the Current Working Directory inside the container
WORKDIR /

# Copy the Go project directory into the container
COPY / .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the backend
CMD ["./main"]
