# syntax=docker/dockerfile:1
# Start from golang v1.13.4 base image to have access to go modules
FROM golang:1.18.1-alpine

# create a working directory
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# copy source from the host to the working directory inside
# the container
COPY . .
RUN go build -o godict ./cmd/*.go

# This container exposes port to the outside world
EXPOSE 8081

CMD ["./godict"]