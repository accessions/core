FROM golang:latest

ENV GOPROXY https://goproxy.io,direct

WORKDIR $GOPATH/src/smoke

COPY . $GOPATH/src/smoke

RUN go build .

ENTRYPOINT ["./smoke"]
