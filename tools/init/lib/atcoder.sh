#!/usr/bin/env bash

SCRIPT_DIR=$(cd $(dirname ${BASH_SOURCE:-$0}); pwd)
source $SCRIPT_DIR/util.sh

ATCODER_URL_REGEX='^https:\/\/atcoder.jp\/contests\/\([a-zA-Z0-9_-]*\)\/tasks\/\([a-zA-Z0-9_-]*\)$'

is_atcoder() {
  local url="$1"
  regex_match "$ATCODER_URL_REGEX" "$url"
}

get_atcoder_project_dir() {
  local url="$1" language="$2"

  local contest_id task_id
  contest_id=$(regex_replace "$ATCODER_URL_REGEX" "$url" "\1")
  task_id="$(regex_replace "$ATCODER_URL_REGEX" "$url" "\2")"

  echo "$language/atcoder/$contest_id/$task_id"
}

create_atcoder_project() {
  local url="$1" language="$2"

  local project_dir
  project_dir="$(get_atcoder_project_dir "$url" "$language")"

  create_project_from_template "$project_dir" "$language" "$url"
}
