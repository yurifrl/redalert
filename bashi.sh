#!/bin/bash
docker run --rm -it --entrypoint=/bin/bash \
  -w /go/src/github.com/jonog/redalert \
  -v "$(pwd):/go/src/github.com/jonog/redalert" \
  -v /var/run/docker.sock:/var/run/docker.sock \
  golang
