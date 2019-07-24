FROM golang:1.12.6-alpine3.9 AS build

ARG GOMAIN=echo-webapp.go

WORKDIR /go/src/webapps
ENV GOPATH /go

RUN apk add -u alpine-sdk

COPY * /go/src/webapps/

RUN go get -d -v ./...
RUN go build -a -o webapp $GOMAIN

FROM alpine:3.9

WORKDIR /go/apps
COPY --from=build /go/src/webapps/webapp .

EXPOSE 3000

ENTRYPOINT ["./webapp"]