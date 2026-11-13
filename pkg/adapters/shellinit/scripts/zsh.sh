pls() {
  local target
  target="$(command pls-cli jump "$@")" || return
  builtin cd -- "$target"
}

_pls_record_visit() {
  command pls-cli add -- "$PWD" 2>/dev/null
}

typeset -ga chpwd_functions
if (( ! ${chpwd_functions[(Ie)_pls_record_visit]} )); then
  chpwd_functions+=(_pls_record_visit)
fi
