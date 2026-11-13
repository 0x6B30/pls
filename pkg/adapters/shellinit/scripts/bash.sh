_pls_cd() {
  local target
  target="$(command pls-cli "$@")" || return
  builtin cd -- "$target"
}

pls() {
  if [[ "${1:-}" == jump ]]; then
    shift
    _pls_cd jump "$@"
    return
  fi
  _pls_cd open "$@"
}

plsj() {
  _pls_cd jump "$@"
}

_pls_record_visit() {
  [[ "${_PLS_LAST_DIR:-}" == "$PWD" ]] && return
  _PLS_LAST_DIR="$PWD"
  command pls-cli add -- "$PWD" 2>/dev/null
}

if [[ "${PROMPT_COMMAND:-}" != *_pls_record_visit* ]]; then
  PROMPT_COMMAND="_pls_record_visit${PROMPT_COMMAND:+;$PROMPT_COMMAND}"
fi
