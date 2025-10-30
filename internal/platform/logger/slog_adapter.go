package logger

import (
	"log/slog"
	"os"
)

type SlogAdapter struct {
	logger *slog.Logger
}

func NewSlogAdapter() *SlogAdapter {
	// You can customize the handler (e.g., JSONHandler) and log level here.
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	return &SlogAdapter{
		logger: slog.New(handler),
	}
}

func (s *SlogAdapter) Info(msg string, args ...any) {
	s.logger.Info(msg, args...)
}

func (s *SlogAdapter) Debug(msg string, args ...any) {
	s.logger.Debug(msg, args...)
}

func (s *SlogAdapter) Warn(msg string, args ...any) {
	s.logger.Warn(msg, args...)
}

func (s *SlogAdapter) Error(err error, msg string, args ...any) {
	allArgs := make([]any, 0, 2+len(args))
	allArgs = append(allArgs, "error", err)
	allArgs = append(allArgs, args...)
	s.logger.Error(msg, allArgs...)
}
