FROM golang:1.13.5 as builder
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

WORKDIR /go/src/admin
COPY . /go/src/admin
RUN mkdir dist \
    && cd service/transfer \
    && CGO_ENABLED=0 go build -a -installsuffix cgo \
    && cp -r /go/src/admin/service/transfer/transfer /go/src/admin/dist

FROM alpine:latest
COPY --from=builder /go/src/admin/dist /usr/bin/server
WORKDIR /usr/bin/server
RUN chmod -R 777 /usr/bin/server

ENTRYPOINT ["./transfer"]
#CMD exec /bin/sh -c "trap : TERM INT; (while true; do sleep 1000; done) & wait"
EXPOSE 80