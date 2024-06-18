package web

import (
	"embed"
	"io/fs"
)

var (
	//go:embed app/dist static templates/*
	root embed.FS

	App       fs.FS
	Static    fs.FS
	Templates fs.FS
)

func init() {
	App, _ = fs.Sub(root, "app/dist")
	Static, _ = fs.Sub(root, "static")
	Templates, _ = fs.Sub(root, "templates")
}
