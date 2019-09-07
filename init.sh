#!/bin/bash -e

if [[ ! -e "bin/golangci-lint" ]]; then
  curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.17.1
else
  echo "[INFO] already installed golangci-lint"
fi
