FROM golang:1.22-bookworm as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the files into the container
COPY . .

# Download dependencies
RUN go mod download

# Build
RUN CGO_ENABLED=1 GOOS=linux go build cmd/main.go

# Start a new stage
FROM debian:bookworm

RUN apt-get update && apt-get install -y \
    ca-certificates \
    sqlite3

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy SQL setup script
COPY ./sql/create_tables.sql /root/create_tables.sql

# Initialize database
RUN sqlite3 db.db < create_tables.sql

# Run app binary
CMD ["./main"]