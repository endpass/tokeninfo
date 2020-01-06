FROM golang:1.12.14-alpine3.11 as builder

RUN apk add --no-cache \
    git \
    curl \
    build-base \
    ca-certificates

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
EXPOSE 8080
ENTRYPOINT ["/tokeninfo"]
