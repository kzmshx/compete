#!/usr/bin/env bash

SCRIPT_DIR=$(cd $(dirname ${BASH_SOURCE:-$0}); pwd)
LIB_DIR="$SCRIPT_DIR/lib"
source $LIB_DIR/atcoder.sh
source $LIB_DIR/language.sh
source $LIB_DIR/template.sh

main() {
  local url
  read -r -p "url: " url

  local language
  PS3="language: "
  select language in "${SUPPORTED_LANGUAGES[@]}"; do
    if [ -n "$language" ]; then
      echo "$language"
      break
    else
      echo "invalid language: $REPLY"
    fi
  done

  echo "$url" "$language"

  if is_atcoder "$url"; then
    create_atcoder_project "$url" "$language"
  else
    echo "unsupported format: $url" && exit 1
  fi
}

main "$@"
