package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
)

const (
	exitOK      = 0
	exitFailure = 1
	exitUsage   = 2
)

const usage = `usage:
  pls-cli open <dir | git-url [clone-dir]>
  pls-cli jump <partial-name>
  pls-cli add <absolute-path>
  pls-cli init <zsh | bash>
`

var errUsage = errors.New("invalid arguments")

type command func(ctx context.Context, args []string) error

func (a *app) commands() map[string]command {
	return map[string]command{
		"open": a.open,
		"jump": a.jump,
		"add":  a.add,
		"init": a.printShellInit,
	}
}

func run(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stderr, usage)
		return exitUsage
	}

	a, err := newApp(stdout, stderr)
	if err != nil {
		fmt.Fprintf(stderr, "pls: %v\n", err)
		return exitFailure
	}

	cmd, ok := a.commands()[args[0]]
	if !ok {
		fmt.Fprintf(stderr, "pls: unknown command %q\n%s", args[0], usage)
		return exitUsage
	}

	return exitCode(cmd(ctx, args[1:]), stderr)
}

func exitCode(err error, stderr io.Writer) int {
	switch {
	case err == nil:
		return exitOK
	case errors.Is(err, flag.ErrHelp):
		return exitUsage
	case errors.Is(err, errUsage):
		fmt.Fprint(stderr, usage)
		return exitUsage
	default:
		fmt.Fprintf(stderr, "pls: %v\n", err)
		return exitFailure
	}
}

func newFlagSet(name string, stderr io.Writer) *flag.FlagSet {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.SetOutput(stderr)
	return flags
}
