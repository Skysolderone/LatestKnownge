package main

import (
	"github.com/pkg/errors"

	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/rs/zerolog"
)

// panic (zerolog.PanicLevel, 5)
// fatal (zerolog.FatalLevel, 4)
// error (zerolog.ErrorLevel, 3)
// warn (zerolog.WarnLevel, 2)
// info (zerolog.InfoLevel, 1)
// debug (zerolog.DebugLevel, 0)
// trace (zerolog.TraceLevel, -1)
// go get -u github.com/rs/zerolog/log
func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// zerolog.ErrorHandler
	// 使用stack一定要打开这个
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log.Print("hello world")
	err := outer()
	log.Info().Msg("test info output")
	log.Error().Stack().Err(err).Msg("")
}

func inter() error {
	err := errors.New("test output stack")
	return err
}

func middle() error {
	err := inter()
	if err != nil {
		return err
	}
	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		return err
	}
	return nil
}
