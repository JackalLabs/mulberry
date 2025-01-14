FROM golang:1.23 AS builder

WORKDIR /app

# Copy the Go module manifests into /app 
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application using the Makefile
RUN make install

# Use a smaller image for the final runtime container
FROM debian:bookworm-slim

# Set up necessary environment variables
ENV MULBERRY_HOME /root/.mulberry

# Install dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Download the libwasmvm.aarch64.so file directly from GitHub
WORKDIR /tmp
RUN wget https://github.com/CosmWasm/wasmvm/releases/download/v2.1.4/libwasmvm.`uname -m`.so \
    -O /usr/local/lib/libwasmvm.`uname -m`.so && \
    chmod 755 /usr/local/lib/libwasmvm.`uname -m`.so

# Update the linker cache
RUN ldconfig /usr/local/lib/

# Set the working directory
WORKDIR /root/

# Copy the binary and configuration files from the builder
COPY --from=builder /go/bin/mulberry /usr/local/bin/mulberry
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