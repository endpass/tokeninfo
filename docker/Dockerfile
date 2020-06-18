FROM golang:1.13.12-alpine3.11 as builder

RUN apk add --no-cache \
    git=2.24.3-r0 \
    curl=7.67.0-r0 \
    build-base=0.5-r1 \
    ca-certificates=20191127-r1

ENV GOPRIVATE=github.com/endpass
ENV GOPROXY=https://proxy.golang.org,direct
ENV GO111MODULE=on
ENV USER=tokeninfo
ENV UID=10001

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

USER tokeninfo
WORKDIR /app
COPY . .
ADD https://github.com/trustwallet/tokens/tree/master/tokens /app/tokens
ADD https://github.com/MyEtherWallet/ethereum-lists/raw/master/dist/tokens/eth/tokens-eth.json /app/tokens-eth.json

RUN go mod download && \
    GOOS=linux GOARCH=amd64 go build -a -tags netgo \
    -ldflags '-s -w -extldflags "-static"' \
    -o bin/tokeninfo

FROM scratch
ENV TOKEN_LIST /data/tokens-eth.json
ENV TOKEN_IMAGE_DIR /data/tokens
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/bin/tokeninfo .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/tokens/ /data/tokens/
COPY --from=builder /app/tokens-eth.json /data/
USER tokeninfo
EXPOSE 8000
ENTRYPOINT ["/tokeninfo"]
