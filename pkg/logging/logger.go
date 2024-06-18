package logging

import (
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
)

type Config struct {
	Level     int    `yaml:"level"`
	Directory string `yaml:"dir"`
}

func NewLogger(cfg *Config) (*slog.Logger, error) {
	const (
		flag int         = os.O_WRONLY | os.O_CREATE | os.O_APPEND
		mode fs.FileMode = 0o0644
	)

	path := filepath.Join(cfg.Directory, "andy.log")
	f, err := os.OpenFile(path, flag, mode)
	if err != nil {
		return nil, err
	}

	w := io.MultiWriter(os.Stdout, f)
	handler := slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.Level(cfg.Level),
	})

	return slog.New(handler), nil
}
