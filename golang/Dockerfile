FROM golang:1.11.10-alpine3.9

ADD /main.go /go/src/workshop/

WORKDIR /go/src/workshop/

RUN go get -d ./... && go build -o k8s .

RUN go install workshop

ENTRYPOINT /go/bin/workshop

EXPOSE 8080
