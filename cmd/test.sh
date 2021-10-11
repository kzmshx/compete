#!/bin/zsh

str_repeat() {
  seq -f "$1" -s "" "$2"
}

bar() {
  echo "$(str_repeat "$1" 100)"
}

build() {
  local src="$1" target="${src%%.*}" cxx="/usr/local/bin/g++-9" cxxflags
  cxxflags=($(cat "$script_dir/../compile_flags.txt"))
  $cxx ${cxxflags[*]} -o "$target" "$src"
  echo "$target"
}

assert() {
  local red='\033[0;31m' green='\033[0;32m' nc='\033[0m'
  local bin="$1" in="$2" out="$3" tmp res

  tmp="$(mktemp)"
  trap "rm -f $tmp" 0

  $bin <"$in" >"$tmp"

  bar "="
  bar "-"
  sdiff "$tmp" "$out"
  res=$?
  bar "-" 100
  if test $res; then
    echo "${green}OK${nc}"
  else
    echo "${red}NG${nc}"
  fi
  bar "=" 100
}

nth_in() { echo "$case_dir/sample-$1.in"; }

nth_out() { echo "$case_dir/sample-$1.out"; }

script_dir="$(dirname "$0")"

src="$1"
src_dir="$(dirname "$src")"
case_dir="$src_dir/test"

bin=$(build "$src")
trap "rm -f $bin" 0

i=1 in="$(nth_in "$i")" out="$(nth_out "$i")"
while [ -f "$in" ] && [ -f "$out" ]; do
  assert "$bin" "$in" "$out"
  i=$((i + 1)) in=$(nth_in $i) out=$(nth_out $i)
done
