FROM golang:1.13.5-alpine3.10 as builder

RUN apk add --no-cache \
    git \
    curl \
    build-base \
    ca-certificates

ENV GOPRIVATE=github.com/machinae
ENV GOPROXY=https://proxy.golang.org,direct
ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN go mod download && \
    GOOS=linux GOARCH=amd64 go build -a -tags netgo \
    -ldflags '-s -w -extldflags "-static"' \
    -o bin/tokeninfo

FROM scratch
ENV TOKEN_LIST /data/tokens-eth.json
ENV TOKEN_IMAGE_DIR /data/tokens
COPY --from=builder /app/bin/tokeninfo .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/tokens/ /data/tokens/
COPY --from=builder /app/tokens-eth.json /data/
HEALTHCHECK --interval=10s --timeout=1m --retries=5 CMD curl http://localhost:8080/health || exit 1
EXPOSE 8080
ENTRYPOINT ["/tokeninfo"]
