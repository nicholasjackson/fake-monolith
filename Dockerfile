FROM golang:latest

COPY . /go/src/github.com/nicholasjackson/fake-monolith

WORKDIR /go/src/github.com/nicholasjackson/fake-monolith

RUN CGO_ENABLED=0 go build -o bin/fake main.go

FROM alpine:latest

RUN apk add -u socat curl

COPY --from=0 /go/src/github.com/shipyard-run/ingress/bin/ingress /usr/local/bin/fake

ENTRYPOINT [ "fake" ]
