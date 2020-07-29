FROM golang:1.14.2-alpine3.11 as buildimage

WORKDIR /src/webserver

COPY webserver.go ./

RUN go build -o /webserver -v webserver.go

FROM alpine:3.11

COPY --from=buildimage /webserver /webserver

ENV VERSION 2.0

ENTRYPOINT ["/webserver"]
