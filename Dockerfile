# # Use the official Golang image as the base image
# FROM golang:1.23.1

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the Go module files and download dependencies
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the rest of the application code
# COPY . .

# # Build the Go application
# RUN go build -o main .

# # Expose the port the app runs on
# EXPOSE 8080

# # Start the application
# CMD ["./main"]


# # Use the official Golang image as the base image
# FROM golang:1.20

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the Go module files and download dependencies
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the rest of the application code
# COPY . .

# # Build the Go application
# RUN go build -o main .

# # Expose the port the app runs on
# EXPOSE 8080

# Start the application
# CMD ["./main"]

FROM golang:1.23

WORKDIR /web

COPY . .

RUN go build -o app .

EXPOSE 80

CMD ["./app"]
