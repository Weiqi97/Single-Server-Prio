# syntax=docker/dockerfile:1

# Choose Alpine for its small footprint.
FROM golang:1.19-alpine

# Set the working directory.
WORKDIR /app

# Download the necessary Go modules for the application.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy over the necessary folders and files.
COPY FLPCP/ ./FLPCP/
COPY IPFE/ ./IPFE/
COPY utilities/ ./utilities/
COPY *.go ./
