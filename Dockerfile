# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/github.com/marioTiara/todolistapp

# Copy the local package files to the container's working directory
COPY . .

# Download and install any required third-party dependencies into the container
RUN go get -d -v ./...

# Build your Go application
RUN go build -o todolistapp ./cmd

# Expose the port on which your application will run
EXPOSE 8080

# Define the command to run your application
CMD ["./todolistapp"]
