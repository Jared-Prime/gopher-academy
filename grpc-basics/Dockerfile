FROM golang:1.10 as builder

WORKDIR /builder

COPY src/go .

RUN go get google.golang.org/grpc
RUN go build example.go

FROM ruby:2.5.0

WORKDIR /app

COPY src .
COPY --from=builder /builder/example example

CMD ["run-example.sh"]
