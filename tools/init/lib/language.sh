#!/usr/bin/env bash

SUPPORTED_LANGUAGES=("cpp" "go" "rs")

is_supported_language() {
  local lang="$1"
  for supported_lang in "${SUPPORTED_LANGUAGES[@]}"; do
    if [ "$lang" = "$supported_lang" ]; then
      return 0
    fi
  done
  return 1
}
