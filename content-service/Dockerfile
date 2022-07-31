# Build stage
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN make build

# Deploy stage
FROM gcr.io/distroless/static AS final

WORKDIR /app

USER nonroot:nonroot

COPY --from=build --chown=nonroot:nonroot /content-service /content-service

EXPOSE 8000

ENTRYPOINT ["/content-service"]