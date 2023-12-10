FROM golang:alpine

ENV DATA_PATH=/data

ENV PORT=555

WORKDIR /go/src/github.com/brtmvdl/antify

COPY . .

CMD go mod tidy && \
  go build && \
  ./antify --data=${DATA_PATH} --port=${PORT}
