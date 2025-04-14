package errs

import (
	"fmt"
	"runtime"

	"github.com/rs/zerolog/log"
)

func callPanic() {
	_, file, line, _ := runtime.Caller(1)
	panic(fmt.Sprintf("panic at %v:%v", file, line))

	log.Info().Msg("this would never be printed!")
}

func callPanicAndRecover() {
	defer func() {
		if v := recover(); v != nil {
			log.Info().Msgf("panic recovered: %v", v)
		}
	}()

	_, file, line, _ := runtime.Caller(1)
	panic(fmt.Sprintf("panic at %v:%v", file, line))

	log.Info().Msg("this would never be printed!")
}
