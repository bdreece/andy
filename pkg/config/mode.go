package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Mode int
const (
    HTTP Mode = iota
    FCGI
)

func (m *Mode) UnmarshalYAML(value *yaml.Node) error {
    var str string
    if err := value.Decode(&str); err != nil {
        return err
    }

    switch str {
    case "http":
        *m = HTTP
    case "fcgi":
        *m = FCGI
    default:
        return fmt.Errorf("invalid mode: %q", str)
    }

    return nil
}
