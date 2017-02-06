FROM golang:alpine3.5
MAINTAINER Stephen Price <stephen@stp5.net>

ENV BUILD_DEPS="git"

WORKDIR ${GOPATH}
RUN apk --update add ${BUILD_DEPS} \
    && go get github.com/steeef/cog-go-rundeck \
    && mkdir -p /bundle \
    && cp ${GOPATH}/bin/cog-go-rundeck /bundle/cog-go-rundeck \
    && cd /bundle \
    && rm -rf /go \
    && apk del ${BUILD_DEPS} \
    && rm -rf /var/cache/apk/*

COPY scripts/* /bundle/
WORKDIR /bundle
RUN chmod +x ./*.sh
