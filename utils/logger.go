package utils

import (
	"os"

	"github.com/ShardulNalegave/freighter/analytics"
	"github.com/rs/zerolog"
)

func NewLogger(a *analytics.Analytics, enable_console_logging bool) zerolog.Logger {
	if enable_console_logging {
		return zerolog.New(zerolog.MultiLevelWriter(
			a.LogFile,
			zerolog.ConsoleWriter{
				Out: os.Stdout,
			},
		)).With().Timestamp().Logger()
	} else {
		return zerolog.New(a.LogFile).With().Timestamp().Logger()
	}
}
