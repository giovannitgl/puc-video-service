BINARY_NAME=content-service

all: build

build:
	go build -o ${BINARY_NAME} main.go

docker-build:
	docker build .

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean:
	go clean