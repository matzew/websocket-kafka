# Multistage build
# docker build -t matzew/ws-kafka

# Envs:

FROM golang:1.13 AS build
MAINTAINER Matthias Wessendorf <matzew@apache.org>


ADD . /src
WORKDIR /src
# RUN go get github.com/Masterminds/glide
RUN CGO=0 go build -o http-kafka ./cmd/bridge/main.go

FROM centos

ENV WEBSOCKET_SERVER wss://localhost/echo
ENV KAFKA_TOPIC=my-topic
ENV KAFKA_BOOTSTRAP_HOST=localhost
ENV KAFKA_BOOTSTRAP_PORT=9092

COPY --from=build /src/http-kafka /http-kafka
ENTRYPOINT ["/http-kafka"]
