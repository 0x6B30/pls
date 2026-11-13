package main

import (
	"context"
	"fmt"
	"os"
)

func (a *app) jump(_ context.Context, args []string) error {
	flags := newFlagSet("jump", a.stderr)
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 1 {
		return errUsage
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	query := flags.Arg(0)
	path, err := a.matcher.Best(query, cwd)
	if err != nil {
		return fmt.Errorf("%w %q", err, query)
	}

	_, err = fmt.Fprintln(a.stdout, path)
	return err
}
