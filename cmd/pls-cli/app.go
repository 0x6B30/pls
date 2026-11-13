package main

import (
	"io"
	"time"

	"github.com/0x6B30/pls/pkg/adapters/shellinit"
	"github.com/0x6B30/pls/pkg/infra/dirs"
	"github.com/0x6B30/pls/pkg/infra/gitcli"
	"github.com/0x6B30/pls/pkg/infra/historystore"
	"github.com/0x6B30/pls/pkg/services/matcher"
	"github.com/0x6B30/pls/pkg/services/resolver"
)

type app struct {
	resolver resolver.Resolver
	visits   historystore.Recorder
	scripts  shellinit.Generator
	stdout   io.Writer
	stderr   io.Writer
}

func newApp(stdout, stderr io.Writer) (*app, error) {
	historyPath, err := historyFilePath()
	if err != nil {
		return nil, err
	}

	store := historystore.New(historyPath)
	directories := dirs.New()

	return &app{
		resolver: resolver.New(resolver.Dependencies{
			Cloner:  gitcli.New(stderr),
			Dirs:    directories,
			Matcher: matcher.New(store, directories, time.Now),
		}),
		visits:  store,
		scripts: shellinit.New(),
		stdout:  stdout,
		stderr:  stderr,
	}, nil
}
