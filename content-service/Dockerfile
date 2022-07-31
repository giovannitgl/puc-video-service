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

RUN make build && ls

# Deploy stage
FROM gcr.io/distroless/static AS final

WORKDIR /app

USER nonroot:nonroot

COPY --from=build --chown=nonroot:nonroot /app/content-service /app/content-service

EXPOSE 8000

ENTRYPOINT ["/content-service"]