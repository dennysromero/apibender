# Start from the latest golang base image

FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Dennys Romero <dennysromero@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o rest-api-mysql .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./rest-api-mysql"]
