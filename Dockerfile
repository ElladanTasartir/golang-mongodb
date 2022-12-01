FROM golang:1.19 as build

WORKDIR /go/bin/golang-mongodb

COPY . .

RUN go mod download

RUN go build -o golang-poc ./cmd/golang-poc/main.go

FROM scratch

COPY --from=build /go/bin/golang-mongodb/golang-poc /golang-poc

EXPOSE 3333

ENTRYPOINT [ "./golang-poc" ]