# Base image
FROM golang:alpine as builder

# Installing required dependencies
RUN apk update && apk add --no-cache git

# Current working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source to destination in container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# New stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates


WORKDIR /root/
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

CMD ["./main"]