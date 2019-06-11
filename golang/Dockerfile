FROM golang:1.11.10-alpine3.9

ENV HOME=/root

ADD /main.go /go/src/workshop/

WORKDIR /go/src/workshop/

RUN go install workshop

ENTRYPOINT /go/bin/workshop

EXPOSE 8080
