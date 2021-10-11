#!/bin/zsh

script_dir="$(dirname "$0")"
source "$script_dir/util.sh"

src_dir="./src"

sub_atcoder() {
  local atcoder_dir="$src_dir/atcoder" url="$1" command="$2" contest_id task_id ans_dir ans_path tmpl_path

  read -r contest_id task_id <<<"$(parse_atcoder_url "$url")"
  ans_dir="$atcoder_dir/$contest_id/$task_id"
  ans_path="$ans_dir/main.cpp"
  tmpl_path="$atcoder_dir/template.cpp"

  atcoder_new() {
    mkdir -p "$ans_dir" && cp "$tmpl_path" "$ans_path" && cd "$ans_dir" && oj d "$url" "$@"
  }

  atcoder_test() {
    cd "$ans_dir" && oj t "$@"
  }

  atcoder_submit() {
    oj s -y "$url" "$ans_path" "$@"
  }

  case "$command" in
  n | new) shift 2 && atcoder_new "$@" && exit 0 ;;
  t | test) shift 2 && atcoder_test "$@" && exit 0 ;;
  s | submit) shift 2 && atcoder_submit "$@" && exit 0 ;;
  *) echo "unknown command: $command" && exit 1 ;;
  esac
}

url="$1"

if is_atcoder "$url"; then
  sub_atcoder "$@" && exit 0
else
  echo "unsupported url: $url" && exit 1
fi
