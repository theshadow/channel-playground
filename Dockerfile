FROM golang:1.13.5-alpine3.11 AS build
LABEL maintainer="Xander Guzman <xander.guzman@xanderguzman.com>"

RUN mkdir -p src/channel-playground
WORKDIR /go/src/channel-playground

COPY . /go/src/channel-playground

RUN go build -o playground .../..

FROM alpine:3.11 as production
LABEL maintainer="Xander Guzman <xander.guzman@xanderguzman.com>"

COPY --from=build /go/src/channel-playground/playground /bin/playground

ENV STATSD_INTERFACE localhost:8125

ENTRYPOINT /bin/playground -statsd $STATSD_INTERFACE -cpu -mem -gc -publish-runtime-stats -metric-prefix channel-playground
