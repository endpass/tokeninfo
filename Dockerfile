FROM golang:1.10.3-alpine
LABEL maintainer="Alex Romanin alexandr@endpass.com"

ENV TOKEN_LIST /opt/data/ethereum-lists/tokens/tokens-eth.json
ENV TOKEN_IMAGE_DIR /opt/data/tokens/images

ADD . /go/src/github.com/endpass/tokeninfo

WORKDIR /go/src/github.com/endpass/tokeninfo

RUN apk update && \
    apk add git curl && \ 
    go get -d -v && \
    go install -v

VOLUME ["/opt/data"]

HEALTHCHECK --interval=10s --timeout=1m --retries=5 CMD curl http://localhost:8000/health || exit 1

ENTRYPOINT ["/go/bin/tokeninfo"]

EXPOSE 8000
