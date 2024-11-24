# Build stage
FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

# Copy go.mod and go.sum files for dependency resolution
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -ldflags="-s -w" -o /app/main main.go

# Final stage
FROM alpine:latest

# Copy necessary files from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ=Asia/Shanghai

# Set the working directory
WORKDIR /app

# Copy the built application
COPY --from=builder /app/main /app/main

# Copy templates
COPY --from=builder /build/templates /app/templates

# Command to run the application
CMD ["./main"]
