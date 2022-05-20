#!/usr/bin/env bash

set -e

/usr/local/bin/consul-template \
  -config "/etc/consul_template.d" \
  -exec   "/app/api"