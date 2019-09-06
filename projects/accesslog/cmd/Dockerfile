FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY accesslog-linux-amd64 /usr/local/bin/accesslog

ENTRYPOINT ["/usr/local/bin/accesslog"]]