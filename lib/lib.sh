#!/bin/zsh

script_dir="$(dirname "$0")"

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

build() {
  local cxx cxxflags src target

  cxx="/usr/local/bin/g++-9"
  cxxflags=($(cat "$script_dir/../compile_flags.txt"))
  src="$(realpath "$1")"
  target="${src%.*}"

  $cxx ${cxxflags[*]} -o "$target" "$src"

  echo "$target"
}

run_tests() {
  local red='\033[0;31m' green='\033[0;32m' nocolor='\033[0m'
  local target="$1" testcase_dir="$2"

  cols() { tput cols; }

  str_repeat() { seq -f "$1" -s "" "$2"; }

  bar() { echo "$(str_repeat "$1" "$(cols)")"; }

  assert() {
    local target="$1" in="$2" out="$3"
    local actual result_diff result_status message

    actual="$(mktemp)"
    trap "rm -f $actual" EXIT
    $target <"$in" >"$actual"

    result_diff="$(sdiff "$out" "$actual" -w "$(cols)")"
    result_status=$?
    if [ $result_status -eq 0 ]; then
      message="${green}OK${nocolor}"
    else
      message="${red}NG${nocolor}"
    fi

    bar "-"
    echo Input:
    echo "$(cat $in)"
    echo
    echo Output:
    echo $result_diff
    echo
    echo Status:
    echo $message
    bar "-"

    return $result_status
  }

  in_file() { echo "$testcase_dir/sample-$1.in"; }

  out_file() { echo "$testcase_dir/sample-$1.out"; }

  local i=1 result_status n_ok=0 n_ng=0
  while [ -f "$(in_file $i)" ] && [ -f "$(out_file $i)" ]; do
    assert "$target" "$(in_file $i)" "$(out_file $i)"
    if [ $? -eq 0 ]; then
      n_ok=$((n_ok + 1))
    else
      n_ng=$((n_ng + 1))
    fi
    i=$((i + 1))
  done

  bar "="
  [ $n_ok -gt 0 ] && echo "${green}${n_ok}${nocolor} passed."
  [ $n_ng -gt 0 ] && echo "${red}${n_ng}${nocolor} failed."
  bar "="
}
