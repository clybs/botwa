FROM golang:1.12.9
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -v
MAINTAINER clybs@yahoo.com