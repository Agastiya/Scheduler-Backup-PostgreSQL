## Stage 1: Build Stage

# Use the latest golang base image for the builder
FROM golang:alpine AS builder

# Set up the working directory
RUN mkdir -p /go/src/service
ADD . /go/src/service
WORKDIR /go/src/service

# Install dependencies
RUN go mod tidy

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -o main main.go
RUN chmod 755 /go/src/service/main

## Stage 2: Run Stage

# Use the latest alpine base image
FROM alpine:latest

# Create a non-root user
RUN addgroup -S kratos && adduser -S kratos -G kratos

# Install necessary packages
RUN apk update && apk upgrade && apk add --no-cache tzdata

# Set up the application directory
RUN mkdir -p /app
WORKDIR /app

# Copy necessary files from the builder stage
COPY --chown=kratos:kratos --from=builder /go/src/service/Environment /go/src/service/Environment
COPY --chown=kratos:kratos --from=builder /go/src/service/main /app

# Switch to the non-root user
USER kratos

# Expose port 7000 to the outside world
EXPOSE 7000

# Command to run the application
CMD /app/main -tag=$environment
