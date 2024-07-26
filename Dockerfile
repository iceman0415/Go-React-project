FROM golang:alpine

RUN apk --update add git bash

WORKDIR /api

ENV GO111MODULE=on

COPY . ./
RUN go build

EXPOSE 8080

CMD ["/api/go-react-example"]