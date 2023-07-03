package logger

import (
	"github.com/caioeverest/fed-its/internal/config"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

// New builds a logger that will be used by the application.
// It will be available to all of the application's dependencies.
func New(cfg *config.Config) *Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	l.WithField("version", cfg.Version)

	return &Logger{l}
}
