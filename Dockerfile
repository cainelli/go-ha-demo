FROM cainelli/go-ha-demo:base

WORKDIR /go/src/github.com/cainelli/go-ha-demo

COPY ./ /go/src/github.com/cainelli/go-ha-demo

ENV GO111MODULE=on

RUN go mod download

RUN go build

EXPOSE 8000

CMD ["./go-ha-demo"]
