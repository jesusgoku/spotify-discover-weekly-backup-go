FROM golang:1.13-alpine as builder

RUN apk add --no-cache build-base

WORKDIR /src

COPY . .
ENV CGO_ENABLED=0
RUN make

FROM alpine:3.11
COPY --from=builder /src/bin /bin
ENTRYPOINT ["/bin/credentials"]

