#!/bin/sh

# Substitute BIN for your bin directory.
# Substitute VERSION for the current released version.
BIN="/usr/local/bin"
VERSION="0.43.2"

install() {
  curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/${1}-$(uname -s)-$(uname -m)" \
    -o "${BIN}/${1}" && \
  chmod +x "${BIN}/${1}"
}

uninstall() {
  BIN="/usr/local/bin" && \
  BINARY_NAME="buf" && \
    rm -f "${BIN}/${1}"
}

uninstall "buf"
install "buf"

uninstall "protoc-gen-buf-breaking"
install "protoc-gen-buf-breaking"

uninstall "protoc-gen-buf-lint"
install "protoc-gen-buf-lint"
