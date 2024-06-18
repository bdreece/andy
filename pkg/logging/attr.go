package logging

import (
	"log/slog"
	"reflect"
)

func Context[T any]() slog.Attr {
	return slog.String("context", reflect.TypeFor[T]().String())
}

func Error(err error) slog.Attr {
	return slog.String("error", err.Error())
}
