#!/usr/bin/env bash
# Update source data and restart server

#Check env vars
[ -z "$TOKEN_LIST" ] && echo "TOKEN_LIST is not set" && exit 1
[ -z "$TOKEN_IMAGE_DIR" ] && echo "TOKEN_IMAGE_DIR is not set" && exit 1

# Kill server if running
pid=$(pgrep tokeninfo|head -n1)
if [[ -n "$pid" ]];then
  kill $pid
  # Wait for exit
  while kill -0 $pid;do sleep 1;done;
fi

cd "${TOKEN_LIST}" && git pull
cd "${TOKEN_IMAGE_DIR}" && git pull

tokeninfo "$@"
