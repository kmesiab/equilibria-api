# Use the official Golang image to create a build artifact.
# This can be used for both building and running the application.
FROM golang:1.22.1 as builder

# Set the working directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum and download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Final stage: Use the same golang image for running the application.
FROM golang:1.22.1

WORKDIR /app

# Copy the compiled application from the previous stage.
COPY --from=builder /app/main .

# Expose port 443 for the application.
EXPOSE 443

# Run the binary.
CMD ["./main"]
