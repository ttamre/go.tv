# Start from the official Go image
FROM golang:1.22 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the source code and Makefile into the container
COPY . .

# Build the Go app using the Makefile
RUN make deps build

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the previous stage
RUN echo ${PWD} && ls -la
COPY --from=builder /build/gotv .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./binary"]
