#!/usr/bin/env bash

command_exists() {
  local command="$1"
  if command -v "$command" >/dev/null 2>&1; then
    return 0
  else
    return 1
  fi
}

regex_match() {
  local pattern="$1" string="$2"
  expr "$string" : "$pattern" >/dev/null
}

regex_replace() {
  local pattern="$1" string="$2" replacement="$3"
  echo "$string" | sed "s/$pattern/$replacement/g"
}

prompt() {
  local message="$1"
  local answer

  read -r -p "$message: " answer
  echo "$answer"
}

prompt_select() {
  local message="$1"
  local options=("${@:2}")
  local option

  PS3="$message: "
  select option in "${options[@]}"; do
    if [ -n "$option" ]; then
      echo "$option"
      return 0
    else
      echo "invalid option: $REPLY" >&2
    fi
  done
  return 1
}
