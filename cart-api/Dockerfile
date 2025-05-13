FROM golang:1.23.1

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
COPY src ./src
COPY *.go ./

RUN go mod tidy
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /cart-api

EXPOSE 8000

# Run
CMD ["/cart-api"]