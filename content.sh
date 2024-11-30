#!/bin/bash

# Directory to search
SEARCH_DIR="${1:-.}"

# Find all .go files in the specified directory
find "$SEARCH_DIR" -type f -name "*.go" | while read -r file; do
  echo "File: $file"
  echo "-----------------------"
  cat "$file"
  echo -e "\n-----------------------\n"
done
