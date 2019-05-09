FROM cainelli/go-ha-demo:base

WORKDIR /go/src/github.com/cainelli/go-ha-demo

COPY ./ /go/src/github.com/cainelli/go-ha-demo

RUN go build

EXPOSE 8000

CMD ["./go-ha-demo"]
