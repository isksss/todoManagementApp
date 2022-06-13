FROM golang:1.18-bullseye

WORKDIR /go/src/app

ENTRYPOINT ["/bin/sh", "-c", "while :; do sleep 10; done"]