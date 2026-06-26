#!/usr/bin/env bash

set -e

if [ $# -eq 0 ]; then
    echo "Usage: $0 <problem-name>"
    exit 1
fi

NAME=$(echo "$*" | tr '[:upper:]' '[:lower:]' | tr ' ' '-')

LAST_NUM=$(
    ls -1d [0-9][0-9][0-9][0-9]-* 2>/dev/null \
    | sed 's/-.*//' \
    | sort -n \
    | tail -1
)

if [ -z "$LAST_NUM" ]; then
    NEXT_NUM=1
else
    NEXT_NUM=$((10#$LAST_NUM + 1))
fi

DIR=$(printf "%04d-%s" "$NEXT_NUM" "$NAME")

mkdir "$DIR"

echo "Created: $DIR"
