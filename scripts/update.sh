#!/usr/bin/env bash
# Update source data and restart server

#Check env vars
[ -z "$TOKEN_LIST" ] && echo "TOKEN_LIST is not set" && exit 1
[ -z "$TOKEN_IMAGE_DIR" ] && echo "TOKEN_IMAGE_DIR is not set" && exit 1

cd "${TOKEN_LIST}" && git pull
cd "${TOKEN_IMAGE_DIR}" && git pull

exit
