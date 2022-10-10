# syntax=docker/dockerfile:1
FROM golang:1.19.1-alpine3.16 AS build_base

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

EXPOSE 8080

RUN cd cmd && go build -o /transaction

CMD [ "/transaction" ]
