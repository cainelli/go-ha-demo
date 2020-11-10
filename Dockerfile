FROM golang:1.15.2-alpine3.12 as builder

WORKDIR /go/src/github.com/getyourguide/simple-http-server/
COPY . /go/src/github.com/getyourguide/simple-http-server/
RUN apk add --no-cache  \
      pkgconfig         \
      bash              \
      gcc               \
      git               \
      musl-dev          \
      make              \
      curl              \
      librdkafka-dev


RUN make build

FROM alpine:3.12

RUN apk add --no-cache bash

RUN mkdir /app/

WORKDIR /app

# https://github.com/aws/aws-sdk-go/issues/2322
RUN apk update && apk add ca-certificates tzdata && rm -rf /var/cache/apk/*

RUN chown nobody:nobody .

USER nobody:nobody

COPY --from=builder /go/src/github.com/getyourguide/simple-http-server/simple-http-server .

EXPOSE 8000

CMD ["./simple-http-server"]
