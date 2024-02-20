package utils

import (
	"github.com/ShardulNalegave/freighter/analytics"
	"github.com/rs/zerolog"
)

func NewLogger(a *analytics.Analytics) zerolog.Logger {
	return zerolog.New(a.LogFile)
}
