#!/bin/bash

set -e

PLUGINS_DIR="plugins"
BINARY_NAME="dpage"

echo "Компиляция сервера..."
go build -o "$BINARY_NAME" dpage.go

if [ ! -d "$PLUGINS_DIR" ]; then
    echo "Папка $PLUGINS_DIR не найдена"
    exit 1
fi

echo "Компиляция плагинов..."

for file in "$PLUGINS_DIR"/*.go; do
    if [ ! -f "$file" ]; then
        echo "Нет '.go' файлов в $PLUGINS_DIR/"
        break
    fi

    if ! grep -q "^package main$" "$file"; then
        continue
    fi

    base_name=$(basename "$file" .go)
    so_file="$PLUGINS_DIR/$base_name.so"

    echo "-> $file <$so_file>"
    go build -buildmode=plugin -o "$so_file" "$file"
done

echo "Сборка завершена успешно!"