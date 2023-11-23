#!/usr/bin/env bash

SCRIPT_DIR=$(cd $(dirname ${BASH_SOURCE:-$0}); pwd)
source $SCRIPT_DIR/util.sh

if ! command_exists jq; then
  echo "jq is not installed" && exit 1
fi

if ! command_exists mustache; then
  echo "mustache is not installed" && exit 1
fi

get_template_dir() {
  local language="$1"
  echo "templates/$language"
}

get_versioned_target_dir() {
  local target_dir="$1"

  next_version=$(find "$target_dir" -maxdepth 1 -type d -name "v*" 2>/dev/null | wc -l | awk '{ print $1 + 1 }')
  echo "$target_dir/v$next_version"
}

create_project_from_template() {
  local target_dir="$1" language="$2" url="$3"

  local template_dir
  template_dir="$(get_template_dir "$language")"

  # Copy template
  target_dir="$(get_versioned_target_dir "$target_dir")"
  mkdir -p $target_dir
  cp $template_dir/* $target_dir

  # Initialize project
  cd "$target_dir" || exit 1

  ## Render template
  local tmp_json
  tmp_json=$(mktemp)
  jq -n --arg url "$url" '$ARGS.named' >"$tmp_json"
  mustache "$tmp_json" "Makefile.mustache" >"Makefile"
  rm "Makefile.mustache"
  rm "$tmp_json"

  ## Initialize project
  make init
  echo "Initialized $target_dir"
  tree --charset unicode
}
