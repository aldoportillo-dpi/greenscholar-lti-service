# Build Stage
# Use a Debian base image for building the application
FROM golang:1.24-bookworm as builder

# Install necessary build tools and libraries
RUN apt-get update && apt-get install -y \
    git \
    gcc \
    libc6-dev \
    ca-certificates

# Set an argument for the GitHub token
ARG GITHUB_TOKEN

# Configure git to use the GitHub token for dependency fetching
RUN git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

# Set the working directory
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# Set execute permission on the binary
RUN chmod +x ./main

# Run Stage
# Use a smaller Debian base image for running the application
FROM debian:bookworm

# Install runtime dependencies, if any
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main /app/main

# Copy the submodules from the builder stage
COPY --from=builder /app/db/migrations /app/db/migrations

# Set the working directory
WORKDIR /app

# This is the port that your application listens to
EXPOSE 80

# Command to run the executable
CMD ["./main"]
