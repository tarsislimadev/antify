FROM golang:1.18

WORKDIR /usr/src/app

COPY app .

RUN go mod init github.com/brtmvdl/antify

RUN go mod tidy

RUN go build main.go

CMD ./main
