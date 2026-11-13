package main

import (
	"context"
	"fmt"
	"os"

	"github.com/0x6B30/pls/pkg/services/resolver"
)

func (a *app) jump(ctx context.Context, args []string) error {
	flags := newFlagSet("jump", a.stderr)
	forceCreate := flags.Bool("n", false, "create the folder even if a visited folder matches")
	if err := flags.Parse(args); err != nil {
		return err
	}

	req, err := jumpRequest(flags.Args(), *forceCreate)
	if err != nil {
		return err
	}

	path, err := a.resolver.Resolve(ctx, req)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(a.stdout, path)
	return err
}

func jumpRequest(args []string, forceCreate bool) (resolver.Request, error) {
	if len(args) < 1 || len(args) > 2 {
		return resolver.Request{}, errUsage
	}

	cwd, err := os.Getwd()
	if err != nil {
		return resolver.Request{}, err
	}

	req := resolver.Request{Target: args[0], Cwd: cwd, ForceCreate: forceCreate}
	if len(args) == 2 {
		req.CloneDir = args[1]
	}
	return req, nil
}
