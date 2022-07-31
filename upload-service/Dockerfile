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

RUN make build

# Deploy stage
FROM gcr.io/distroless/static AS final

WORKDIR /app

USER nonroot:nonroot

COPY --from=build --chown=nonroot:nonroot /app/upload-service /app/upload-service

EXPOSE 8000

ENTRYPOINT ["/upload-service"]