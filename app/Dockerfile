# Use a specific version of Golang for a more consistent and stable build environment
FROM golang:1.21.3 AS builder

# Set the working directory inside the container
WORKDIR /app

ARG AWS_ACCESS_KEY_ID
ARG AWS_SECRET_ACCESS_KEY


ENV AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID
ENV AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY

# Copy go.mod and go.sum files first; they change less often than the source code
COPY go.mod go.sum ./

# Download dependencies before copying the source code. 
# This leverages Docker cache to speed up subsequent builds if dependencies haven't changed.
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application. The main file is located in the cmd directory.
# Consider adding flags to create a smaller, optimized binary for a production build.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/...

# Use a smaller base image to run the application
# This reduces the size of the final Docker image and the attack surface area.
FROM alpine:latest  

# Install CA certificates for HTTPS calls
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

#Delete old static and view files
RUN rm -rf /root/static /root/views

# Copy static assets from the builder stage

COPY ./static/styles.css /root/static/styles.css
COPY ./static/script.js /root/static/script.js
COPY ./static/components.js /root/static/components.js
COPY ./views /root/views

# Set the binary as the entrypoint of the container
CMD ["./main"]
