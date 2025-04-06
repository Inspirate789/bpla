package app

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func NewLogger(level slog.Level, addSource, colorOutput bool) *slog.Logger {
	return slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		AddSource: addSource,
		Level:     level,
		NoColor:   !colorOutput,
	}))
}
