# Stage 1: Build the application
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the binary
COPY . .
RUN go build -o fadak .

# Stage 2: Create a lightweight image to run the application
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/fadak .
EXPOSE 8080
CMD ["./fadak"]
