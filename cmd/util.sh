#!/bin/zsh

lower() { cat - | tr '[:upper:]' '[:lower:]'; }

regex_match() {
  local pattern="$1" string="$2"
  expr "$string" : "$pattern" >/dev/null
}

regex_replace() {
  local pattern="$1" replacement="$2" string="$3"
  echo "$string" | sed "s/$pattern/$replacement/"
}

atcoder_url_pattern() { echo '^https:\/\/atcoder.jp\/contests\/\([a-zA-Z0-9_-]*\)\/tasks\/\([a-zA-Z0-9_-]*\)$'; }

is_atcoder() {
  local url="$1"
  regex_match "$(atcoder_url_pattern)" "$url"
}

parse_atcoder_url() {
  local url="$1" contest_id task_id
  contest_id="$(regex_replace "$(atcoder_url_pattern)" "\1" "$url" | lower)"
  task_id="$(regex_replace "$(atcoder_url_pattern)" "\2" "$url" | lower)"
  echo "$contest_id" "$task_id"
}
