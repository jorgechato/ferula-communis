FROM golang:latest

RUN go get github.com/jorgechato/ferula-communis

ENTRYPOINT /go/bin/ferula-communis

EXPOSE 8080
