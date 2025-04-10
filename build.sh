#!/bin/bash

SCRIPT_PWD="$(realpath "${BASH_SOURCE[0]}")"
SCRIPT_DIR="$(dirname "${SCRIPT_PWD}")"

cd "$SCRIPT_DIR" || exit 1

if [[ "$#" -eq "2" ]]; then
    BIN_PATH="$1"
    case "$2" in
    payment | payments)
        GOMAIN_PATH="${SCRIPT_DIR}/cmd/payments/main.go"
        ;;
    decrypt | crypto | encrypt)
        GOMAIN_PATH="${SCRIPT_DIR}/cmd/decrypt/main.go"
        ;;
    *)
        echo 'invalid second parameter'
        exit 1
        ;;
    esac
else
    echo 'invalid amount of parameters'
    exit 1
fi

# automagically build go executable ONLY if go source code was updated in the meanwhile
if [[ -f "$BIN_PATH" ]]; then
    latest_go_file=$(find "$SCRIPT_DIR" -type f -name "*.go" -printf "%T@ %p\n" | sort -nr | head -n1 | awk '{print $2}')
    other_file="$BIN_PATH"

    if [[ $(stat -c %Y "$latest_go_file") -gt $(stat -c %Y "$other_file") ]]; then
        echo 'compiling program...'
        [[ -f "$BIN_PATH" ]] && rm "$BIN_PATH"
        go build -o "$BIN_PATH" "$GOMAIN_PATH"
    fi
else
    go build -o "$BIN_PATH" "$GOMAIN_PATH"
fi
