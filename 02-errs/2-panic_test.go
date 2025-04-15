package errs

import (
	"testing"
)

func Test_callPanic(t *testing.T) {
	callPanic()
}

func Test_callPanicAndRecover(t *testing.T) {
	callPanicAndRecover()
}
