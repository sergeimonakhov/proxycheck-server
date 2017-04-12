#!/bin/sh
if [ $# -eq 0 ]; then
  proxycheck-server
else
  exec "$@"
fi
