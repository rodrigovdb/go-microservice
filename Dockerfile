FROM golang:alpine

ADD . /go/src/app

RUN go install -i app

ENTRYPOINT /go/bin/app

EXPOSE 8080
