
FROM golang

ADD /main.go /go/src/workshop/

WORKDIR /go/src/workshop/

RUN go get -v .

RUN go install workshop

ENTRYPOINT /go/bin/workshop

EXPOSE 8080
