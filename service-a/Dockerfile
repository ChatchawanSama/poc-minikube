# Start from a Golang base image
FROM golang:1.21

WORKDIR /app

COPY . .

# Build the Go app
RUN go build -o main .

# Set the entrypoint
ENTRYPOINT ["/app/main"]

EXPOSE 8080