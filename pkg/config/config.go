package config

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"

	"github.com/bdreece/andy/internal/content"
	"github.com/bdreece/andy/pkg/database"
	"github.com/bdreece/andy/pkg/logging"
	"github.com/bdreece/andy/pkg/session"
)

type Root struct {
	Mode Mode
	Port int
}

func (cfg Root) Addr() string {
	if cfg.Port == 0 {
		switch cfg.Mode {
		case HTTP:
			cfg.Port = 80
		case FCGI:
			cfg.Port = 9000
		default:
			panic(fmt.Errorf("invalid mode: %q", cfg.Mode))
		}
	}

	return fmt.Sprintf(":%d", cfg.Port)
}

type Config struct {
	*Root    `yaml:",inline"`
	Content  *content.Config  `yaml:"content"`
	Database *database.Config `yaml:"database"`
	Logging  *logging.Config  `yaml:"logging"`
	Session  *session.Config  `yaml:"session"`
}

func Decode(r io.Reader) (*Config, error) {
	cfg := new(Config)
	if err := yaml.NewDecoder(r).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
