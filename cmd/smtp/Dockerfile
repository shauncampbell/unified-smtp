FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/shauncampbell/unified-smtp/
COPY . .

RUN go build -o /go/bin/smtp ./cmd/smtp

FROM alpine:3.12

COPY --from=builder /go/bin/imap /go/bin/smtp
LABEL maintainer="Shaun Campbell <docker@shaun.scot>"

VOLUME /etc/smtp.yaml

ENTRYPOINT ["./go/bin/smtp", "-f", "/etc/smtp.yaml"]