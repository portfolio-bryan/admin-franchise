FROM golang:1.20.4-alpine as builder

ENV ENVIRONMENT=dev
ENV SERVER_PORT=8080

WORKDIR /app

COPY go.* .

RUN go mod download

RUN go mod verify

COPY . .

RUN go build -v -o /app server.go

EXPOSE $SERVER_PORT

CMD ["/app/main"]
