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
  local makefile_template="Makefile.mustache"
  local makefile="Makefile"
  local view
  view=$(mktemp)

  jq -n '$ARGS.named' >"$view" \
    --arg "url" "$url"

  mustache "$view" "$makefile_template" >"$makefile"

  rm "$makefile_template"
  rm "$view"

  ## Initialize project
  make init
  echo "Initialized $target_dir"
  tree --charset unicode

  ## Copy target dir to clipboard
  echo "$target_dir" | pbcopy
}
