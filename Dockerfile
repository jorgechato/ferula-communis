FROM golang:latest

ADD . /go/src/github.com/jorgechato/ferula-communis

RUN go install github.com/jorgechato/ferula-communis

ENTRYPOINT /go/bin/ferula-communis

EXPOSE 8080
