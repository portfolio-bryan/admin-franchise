.PHONY:

build-graphql-model:
	go run github.com/99designs/gqlgen generate

install:
	go get ./...

prepare: install migrate-up

start:
	go run cmd/main.go

test:
	go test -v -coverpkg=./... -coverprofile=coverage.out ./app/tests/...

cover:
	go tool cover -func=coverage.out

lint:
	golangci-lint run ./...

build:
	go build -o main ./cmd/main.go

docker-build:
	docker build -t backoffice-business-app:v1 -f docker/Dockerfile .