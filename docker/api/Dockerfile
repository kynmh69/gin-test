FROM golang:1 AS builder

WORKDIR /go/app
COPY ./ ./

RUN go install github.com/cosmtrek/air@latest

CMD [ "air", "-c", ".air.toml" ]
