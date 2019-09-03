FROM golang:latest as builder
LABEL maintainer="pav <pavegy@gmail.com>"

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/github.com/paveg/goura
COPY . .
RUN make build

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/paveg/goura/bin/goura /goura

CMD ["/goura"]
