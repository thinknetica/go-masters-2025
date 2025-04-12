package interfaces

import (
	"math/rand"
	"testing"
)

func Test_uploadFile(t *testing.T) {
	var u uploader

	if rand.Intn(10) < 5 {
		u = &s3uploader{}
	} else {
		u = &azureuploader{}
	}

	uri, err := uploadFile(u, "3-duck_type.go")
	if err != nil {
		t.Fatalf("uploadFile() error = %v", err)
	}

	t.Log(uri)
}
