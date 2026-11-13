package main

import (
	"context"
	"time"
)

func (a *app) add(_ context.Context, args []string) error {
	flags := newFlagSet("add", a.stderr)
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 1 {
		return errUsage
	}

	return a.visits.Record(flags.Arg(0), time.Now())
}
