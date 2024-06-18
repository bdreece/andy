package database

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func applyConfig(ctx context.Context, db DBTX, cfg *Config) error {
    f, err := os.Open(cfg.SchemaPath)
    if err != nil {
        const tpl string = "database: failed to open schema file: %q"
        return errors.Join(fmt.Errorf(tpl, cfg.SchemaPath), err)
    }

    buf := new(bytes.Buffer)
    if _, err = io.Copy(buf, f); err != nil {
        const tpl string = "database: failed to read schema file: %q"
        return errors.Join(fmt.Errorf(tpl, cfg.SchemaPath), err)
    }

    statements := strings.Split(buf.String(), ";")
    for _, statement := range statements {
        if _, err = db.ExecContext(ctx, statement); err != nil {
            const tpl string = "database: failed to apply schema statement: %q"
            return errors.Join(fmt.Errorf(tpl, statement), err)
        }
    }

    return nil
}
