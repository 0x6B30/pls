package main

import (
	"context"
	"fmt"
)

func (a *app) printShellInit(_ context.Context, args []string) error {
	if len(args) != 1 {
		return errUsage
	}

	script, err := a.scripts.Script(args[0])
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(a.stdout, script)
	return err
}
