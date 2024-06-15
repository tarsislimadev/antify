FROM golang:alpine

ENV PORT=80

WORKDIR /usr/src/app

COPY . .

RUN go mod init github.com/brtmvdl/antify

RUN go mod tidy

CMD go run main.go
