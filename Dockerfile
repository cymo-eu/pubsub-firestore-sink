FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/cymo-eu/pubsub-firestore-sink/
COPY . .

RUN  CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' -tags timetzdata

FROM scratch
COPY --from=builder /go/bin/pubsub-firestore-sink /pubsub-firestore-sink
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/pubsub-firestore-sink"]