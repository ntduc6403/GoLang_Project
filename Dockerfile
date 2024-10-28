# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
##

# Use Alpine Go base image for the build stage
FROM golang:1.20.2-alpine as builder

# Create a working directory inside the image
WORKDIR /app

# Copy go.mod and go.sum (if available) first to cache module downloads
# Copy go.mod file into /app
COPY go.mod .  

# Download dependencies before copying source code to cache the downloaded modules
RUN go mod download

# Copy the entire source code into the container (all Go files and subdirectories)
# Copy the current directory contents into /app within the container
COPY . .  

# Compile the application with options to create the binary file in the /app directory
# Compile the main.go file into a binary named /api
RUN go build -o /api ./cmd/api/main.go

##
## STEP 2 - DEPLOY
##

# Use scratch, an empty image for production
FROM scratch

# Copy the binary file from the builder stage to the empty container
# Copy the /api binary file into the empty container
COPY --from=builder /api /api  

# Specify the binary to run when the container starts
ENTRYPOINT ["/api"]
