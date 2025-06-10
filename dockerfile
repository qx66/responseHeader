FROM registry.cn-hangzhou.aliyuncs.com/startops-base/golang-builder:1.20 AS builder

WORKDIR /go/src
ADD . /go/src

ENV GOPROXY=https://goproxy.cn
ENV CGO_ENABLED=0
RUN go build main.go

#FROM docker.io/library/busybox:stable-glibc
FROM registry.cn-hangzhou.aliyuncs.com/startops-base/debian-runtime:11.7-slim

COPY --from=builder /go/src/bin/main /app

WORKDIR /app

CMD ["/app/main"]