#!/bin/bash

set -e

PLUGINS_DIR="plugins"

so_files=$(find "$PLUGINS_DIR" -name "*.so" -type f 2>/dev/null)
if [ -n "$so_files" ]; then
    echo "Удаляю старые плагины:"
    echo "$so_files" | while read -r f; do
        echo "-> $f"
        rm -f "$f"
    done
fi