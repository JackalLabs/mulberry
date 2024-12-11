FROM golang:1.23 AS builder

WORKDIR /app

# Copy the Go module manifests into /app 
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Mulberry server binary
RUN go build -o mulberry main.go

# Use a smaller image for the final runtime container
FROM debian:bullseye-slim

# Set up necessary environment variables
ENV MULBERRY_HOME /root/.mulberry

# Install dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /root/

# Copy the binary and configuration files from the builder
COPY --from=builder /app/mulberry /usr/local/bin/mulberry
COPY --from=builder /app/config /root/.mulberry

# Expose ports - not sure we need all of these 
EXPOSE 26657
EXPOSE 26656
EXPOSE 26658
EXPOSE 1317
EXPOSE 6060
EXPOSE 9090

# Command to initialize the system and start the service
CMD ["sh", "-c", "mulberry wallet address && mulberry start"]