FROM golang:1.13.6 AS builder
LABEL maintainer="psvmcc@gmail.com"

ADD . /src
WORKDIR /src

ENV CGO_ENABLED=0
RUN make build

FROM alpine:3.12.0

USER nobody

COPY --from=builder /src/typidor-exporter /typidor-exporter

EXPOSE     9824

ENTRYPOINT ["/typidor-exporter"]
