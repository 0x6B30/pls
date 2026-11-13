pls() {
  local target
  target="$(command pls-cli jump "$@")" || return
  builtin cd -- "$target"
}

_pls_record_visit() {
  [[ "${_PLS_LAST_DIR:-}" == "$PWD" ]] && return
  _PLS_LAST_DIR="$PWD"
  command pls-cli add -- "$PWD" 2>/dev/null
}

if [[ "${PROMPT_COMMAND:-}" != *_pls_record_visit* ]]; then
  PROMPT_COMMAND="_pls_record_visit${PROMPT_COMMAND:+;$PROMPT_COMMAND}"
fi
