package interfaces

import (
	"testing"
)

func Test_anyType(t *testing.T) {
	anyType()
}

func Test_anyParam(t *testing.T) {
	anyParam(1, "2", func() {}, nil)
}
