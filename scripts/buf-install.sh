#!/bin/bash

# set -x
GOPATH=`go env GOPATH` && \
BIN="/bin" && \
BUF_BIN=$GOPATH$BIN && \
VERSION="1.0.0-rc12" && \
curl -o "${BUF_BIN}/buf" -sSL "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" && \
chmod +x "${BUF_BIN}/buf"