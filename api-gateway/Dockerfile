FROM golang:1.23.1

# Set working directory
WORKDIR /app

# Copy sources
COPY . .

# Download dependencies
RUN go mod tidy
RUN go mod download

# Build the gateway binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-gateway

# Expose gateway port
EXPOSE 8002

# Run the gateway
CMD ["/api-gateway"]