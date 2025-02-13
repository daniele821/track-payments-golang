#!/bin/bash

SCRIPT_PWD="$(realpath "${BASH_SOURCE[0]}")"
SCRIPT_DIR="$(dirname "${SCRIPT_PWD}")"

GOMAIN_PATH="${SCRIPT_DIR}/cmd/payments/main.go"
BIN_PATH="${SCRIPT_DIR}/data/.payments"

[[ -v BUILD ]] && rm "$BIN_PATH"

[[ -f "$BIN_PATH" ]] || go build -o "$BIN_PATH" "$GOMAIN_PATH"

"$BIN_PATH" "$@"
