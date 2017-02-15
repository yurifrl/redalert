#!/bin/bash
docker run --rm \
  -v "$(pwd):/src" \
  -v /var/run/docker.sock:/var/run/docker.sock \
  centurylink/golang-builder yurifl/redalert
