FROM golang:1.19 as build

WORKDIR /go/bin/golang-mongodb

COPY . .

RUN go mod download

WORKDIR /go/bin/golang-mongodb/cmd/golang-poc

RUN go build -o /golang-poc

FROM scratch

COPY --from=build /golang-poc /golang-poc

EXPOSE 3333

RUN [ "/golang-poc" ]