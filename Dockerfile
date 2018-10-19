FROM centos:7
MAINTAINER Matthias Wessendorf <matzew@apache.org>
ARG BINARY=./ws_kafka

COPY ${BINARY} /opt/ws_kafka
ENTRYPOINT ["/opt/ws_kafka"]
