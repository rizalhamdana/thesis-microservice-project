# Start from golang base image
FROM golang:1.14.0-alpine3.11 as builder


# Maintainer info
LABEL maintainer="Rizal Hamdan <ari.gusti12@gmail.com>"

# Install git
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /app

# Copy go.mod and g.sum files 
COPY go.mod go.sum ./

# Download all dependencies 
RUN go mod download

# Copy the source from current directory to the working Directory inside container
COPY . . 

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src/main.go


# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

# Expose port 8080 to the outside world
EXPOSE 8082

#Command to run the executable
CMD ["./main"]