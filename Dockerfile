FROM golang:1.12-stretch as builder
ENV GO111MODULE=on
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    git curl
WORKDIR /app
COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -a -tags netgo \
    -ldflags '-w -extldflags "-static"' \
    -o bin/tokeninfo

FROM alpine:3.9 as certs
RUN apk add --no-cache ca-certificates

FROM scratch
LABEL maintainer="Alex Romanin alexandr@endpass.com"

ENV PORT 8000
ENV TOKEN_LIST /opt/data/ethereum-lists/tokens/tokens-eth.json
ENV TOKEN_IMAGE_DIR /opt/data/tokens/images
COPY --from=builder /app/bin/tokeninfo .
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
VOLUME ["/opt/data"]
HEALTHCHECK --interval=10s --timeout=1m --retries=5 CMD curl http://localhost:8000/health || exit 1
ENTRYPOINT ["/tokeninfo"]
EXPOSE ${PORT}
