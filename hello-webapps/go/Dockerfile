FROM golang:1.12.6-alpine3.9 AS build

ARG GOMAIN=echo-webapp.go

WORKDIR /go/src/webapps
ENV GOPATH /go

RUN apk -U add ca-certificates build-base git && \
    rm -r /var/cache/apk/*

COPY . /go/src/webapps

RUN go get -d -v ./...
RUN go build -a -o webapp $GOMAIN

FROM alpine:3.9

WORKDIR /go/apps
COPY --from=build /go/src/webapps/webapp .

EXPOSE 3000

ENTRYPOINT ["./webapp"]