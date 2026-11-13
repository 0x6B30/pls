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
  command pls-cli add -- "$PWD" 2>/dev/null
}

typeset -ga chpwd_functions
if (( ! ${chpwd_functions[(Ie)_pls_record_visit]} )); then
  chpwd_functions+=(_pls_record_visit)
fi
