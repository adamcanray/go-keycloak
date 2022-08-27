# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

ADD . /app

RUN go build -o /go-keycloak

EXPOSE 8080

CMD [ "/go-keycloak" ]