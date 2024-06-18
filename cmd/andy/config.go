package main

import (
	"os"

	"go.uber.org/fx"

	"github.com/bdreece/andy/internal/content"
	"github.com/bdreece/andy/pkg/cli"
	"github.com/bdreece/andy/pkg/config"
	"github.com/bdreece/andy/pkg/database"
	"github.com/bdreece/andy/pkg/logging"
	"github.com/bdreece/andy/pkg/session"
)

func loadConfig(args cli.Args) (out struct {
	fx.Out

	Root     *config.Root
	Content  *content.Config
	Database *database.Config
	Logging  *logging.Config
	Session  *session.Config
}, err error) {
	f, err := os.Open(args.ConfigPath)
	if err != nil {
		return
	}

	cfg, err := config.Decode(f)
	if err != nil {
		return
	}

	out.Root = cfg.Root
    out.Content = cfg.Content
	out.Database = cfg.Database
	out.Logging = cfg.Logging
	out.Session = cfg.Session

	return
}
