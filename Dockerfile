FROM golang:latest

WORKDIR /go/src/app

COPY . .

WORKDIR /go/src/app/cmd/golang-poc

RUN go get -d -v .

RUN go install -v .

RUN go build -o ../../golang-mongodb

EXPOSE 8080

WORKDIR /go/src/app

RUN ["./golang-mongodb"]