# Use the official Go image as a parent image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace.
COPY . .

# Install dependencies (if you have any)
RUN go mod download

# Build your application
RUN go build -o go-app .

# Make port 8080 available to the world outside this container
EXPOSE 3000

# Run the binary program produced by `go install`
CMD ["./go-app"]
