# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.12.8

# Add Maintainer Info
LABEL maintainer="Nuryanto <nurynaatofattih@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app
WORKDIR /app/runtime
# Build Args
ARG LOG_DIR=/app/runtime/logs


# Create Log Directorytail
RUN mkdir -p ${LOG_DIR}

# Environment Variables
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log 

# Copy go mod and sum files
# COPY go.mod go.sum ./
COPY go.mod go.sum config.json ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# This container exposes port 8080 to the outside world
EXPOSE 8084

# Declare volumes to mount
VOLUME [${LOG_DIR}]

# Run the binary program produced by `go install`
CMD ["./main"]