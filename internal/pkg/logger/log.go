package log

import (
	"log/slog"
	"os"
)

func NewLogger() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func NewTextLogger() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
}

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}
