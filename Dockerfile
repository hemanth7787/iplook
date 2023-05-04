# Use a golang runtime image as the base image
FROM golang:alpine AS build

# Set the working directory to /app
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go binary and statically link libraries
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/server
# RUN go build -ldflags '-extldflags "-static"' -o /go/bin/server

# Create a new container and copy the server binary into it
FROM alpine:latest
COPY --from=build /go/bin/server /go/bin/server

# Install curl to be able to test the HTTP server
RUN apk add --no-cache curl

# Expose port 8080 for the HTTP server
EXPOSE 8080

# Run the Go HTTP server when the container starts
ENTRYPOINT ["/go/bin/server"]