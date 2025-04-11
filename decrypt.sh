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

# automagically build go executable ONLY if go source code was updated in the meanwhile
if [[ -f "$BIN_PATH" ]]; then
    latest_go_file=$(find "$SCRIPT_DIR" -type f -name "*.go" -printf "%T@ %p\n" | sort -nr | head -n1 | awk '{print $2}')
    other_file="$BIN_PATH"

    if [[ $(stat -c %Y "$latest_go_file") -gt $(stat -c %Y "$other_file") ]]; then
        echo 'compiling program (decrypt) ...'
        [[ -f "$BIN_PATH" ]] && rm "$BIN_PATH"
        go build -o "$BIN_PATH" "$GOMAIN_PATH"
    fi
else
    echo 'compiling program (decrypt) ...'
    go build -o "$BIN_PATH" "$GOMAIN_PATH"
fi
