FROM golang:latest as builder

MAINTAINER Huiming, Zhang
COPY ./ /go/src
WORKDIR /go/src
RUN cp /go/src/vendor/github.com /go/src -r
RUN go build -o ops-agent

FROM alpine:latest

COPY --from=builder /go/src/ops-agent /
RUN chmod 777 /ops-agent
CMD "/ops-agent"
