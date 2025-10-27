# Build stage
FROM golang:1.24-alpine AS build

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/aa/main.go

# Run stage
FROM alpine:3.18
# FROM linuxserver/ffmpeg:amd64-8.0-cli-ls44

WORKDIR /app

# Copy binary from build stage
COPY --from=build /app/main .

# For uploads files
RUN mkdir uploads

# Expose port
EXPOSE 9080

# Run the app
# ENTRYPOINT [ "/bin/bash" ]
CMD ["/app/main"]