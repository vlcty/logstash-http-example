FROM golang:rc-alpine AS builder
RUN mkdir -p /go/src/github.com/vlcty/logstash-http-plugin
WORKDIR src/github.com/vlcty/logstash-http-plugin
COPY main.go .
ENTRYPOINT go run main.go
