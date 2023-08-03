FROM golang:latest

WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /player-score-management

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./player-score-management"]
