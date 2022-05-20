#!/usr/bin/env sh

set -e

/usr/local/bin/consul-template \
  -config "/etc/consul-template.d" \
  -exec   "/app/api"