# Build stage
FROM golang:1.18-alpine AS build

WORKDIR /app

RUN apk update &&\
    apk upgrade &&\
    apk add --no-cache make

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN make build && ls -l

# Deploy stage
FROM alpine:3.16.2 as final

WORKDIR /app

RUN addgroup -S cetacean && adduser -S mobydick -G cetacean

COPY --from=build --chown=mobydick:cetacean /app/upload-service /app/upload-service

USER mobydick

ENTRYPOINT ["/app/upload-service"]