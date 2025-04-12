package interfaces

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = zerolog.New(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    true,
			TimeFormat: "2006-01-02 15:04:05",
		},
	).With().Timestamp().Logger().With().Caller().Logger()
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type T struct{}

func (t *T) Read(p []byte) (n int, err error) {
	return 0, nil
}
func (t *T) Write(p []byte) (n int, err error) {
	return 0, nil
}

// Базовая механика работы с интерфейсами.
func basics() {
	var t *T
	var r Reader
	if r == nil {
		log.Info().Msg("r is nil")
	} else {
		log.Info().Msg("r is not nil")
	}

	r = t
	log.Info().Msg("r = t")

	if t == nil {
		log.Info().Msg("t is nil")
	} else {
		log.Info().Msg("t is not nil")
	}

	if r == nil {
		log.Info().Msg("r is nil")
	} else {
		log.Info().Msg("r is not nil")
	}

	r.Read(nil)

	// Type assertion (interface)
	rw, ok := r.(ReadWriter)
	if ok {
		log.Info().Msg("r is ReadWriter")
		rw.Write(nil)
	}

	// Type assertion (type)
	val, ok := rw.(*T)
	if ok {
		log.Info().Msg("rw is *T")
		val.Read(nil)
	}

	// Type switch
	log.Info().Msg("Type switch")
	switch v := rw.(type) {
	case nil:
		log.Info().Msg("rw is nil")
	case Reader:
		log.Info().Msg("rw is Reader")
		v.Read(nil)
	case Writer:
		log.Info().Msg("rw is Writer")
	case ReadWriter:
		log.Info().Msg("rw is ReadWriter")
	default:
		log.Info().Msg("rw is unknown type")
	}
}
