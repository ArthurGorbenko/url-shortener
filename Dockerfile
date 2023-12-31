# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-crud-rest-api

CMD ["/golang-crud-rest-api"]