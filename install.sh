#!/bin/sh
set -eu

MODULE="github.com/0x6B30/pls/cmd/pls-cli"
VERSION="${PLS_VERSION:-latest}"

log() { printf 'pls: %s\n' "$*"; }
fail() { printf 'pls: %s\n' "$*" >&2; exit 1; }

require_go() {
  command -v go >/dev/null 2>&1 || fail "Go is required, get it from https://go.dev/dl/"
}

detect_shell() {
  shell_name="${1:-$(basename "${SHELL:-}")}"
  case "$shell_name" in
    zsh | bash) printf '%s' "$shell_name" ;;
    *) fail "unsupported shell '$shell_name', run: ./install.sh zsh|bash" ;;
  esac
}

rc_file_for() {
  case "$1" in
    zsh) printf '%s/.zshrc' "${ZDOTDIR:-$HOME}" ;;
    bash) printf '%s/.bashrc' "$HOME" ;;
  esac
}

go_bin_dir() {
  bin_dir="$(go env GOBIN)"
  [ -n "$bin_dir" ] || bin_dir="$(go env GOPATH)/bin"
  printf '%s' "$bin_dir"
}

install_binary() {
  log "installing $MODULE@$VERSION"
  go install "$MODULE@$VERSION"
}

hook_rc_file() {
  shell_name="$1"
  rc_file="$2"
  bin_dir="$3"

  if [ -f "$rc_file" ] && grep -q 'pls-cli init' "$rc_file"; then
    log "$rc_file already loads pls, leaving it untouched"
    return
  fi

  {
    printf '\n# pls\n'
    case ":$PATH:" in
      *":$bin_dir:"*) ;;
      *) printf 'export PATH="$PATH:%s"\n' "$bin_dir" ;;
    esac
    printf 'eval "$(pls-cli init %s)"\n' "$shell_name"
  } >>"$rc_file"
  log "added pls to $rc_file"
}

main() {
  require_go
  shell_name="$(detect_shell "${1:-}")"
  rc_file="$(rc_file_for "$shell_name")"

  install_binary
  hook_rc_file "$shell_name" "$rc_file" "$(go_bin_dir)"
  log "done, open a new terminal or run: source $rc_file"
}

main "$@"
