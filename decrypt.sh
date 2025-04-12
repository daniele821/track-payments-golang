#!/bin/bash

SCRIPT_PWD="$(realpath "${BASH_SOURCE[0]}")"
SCRIPT_DIR="$(dirname "${SCRIPT_PWD}")"

cd "$SCRIPT_DIR" || exit 1

GOMAIN_PATH="${SCRIPT_DIR}/cmd/decrypt/main.go"

if [[ "$#" -eq "1" ]]; then
    BIN_PATH="$1"
else
    echo 'invalid amount of parameters'
    exit 1
fi

# build executable binary
go build -o "$BIN_PATH" "$GOMAIN_PATH"
