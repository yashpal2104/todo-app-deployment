# --- Stage 1: The Builder ---
# Use the official Go image to build our application.
# This image contains all the necessary compilers and tools.
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the dependency files first. This leverages Docker's layer caching.
# If these files don't change, Docker won't re-download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application.
# CGO_ENABLED=0 creates a statically linked binary (no external dependencies).
# -o /go-todo-app specifies the output path for the compiled binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-todo-app ./main.go

# --- Stage 2: The Final Image ---
# Use a minimal base image like Alpine Linux.
# This image is much smaller than the Go image, leading to a tiny final image.
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy ONLY the compiled binary from the 'builder' stage.
COPY --from=builder /go-todo-app .

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the application when the container starts.
CMD ["./go-todo-app"]
