# Use the Golang Alpine image for building the application
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /myapp

# Use Alpine as the final base image
FROM alpine:3.18

# Install libc6-compat for glibc compatibility
RUN apk add --no-cache libc6-compat

# Copy the binary from the builder image
COPY --from=builder /myapp /myapp
COPY ./.env.example ./.env

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["/bin/sh", "-c", "sleep 10 && /myapp"]
