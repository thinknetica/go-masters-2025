package interfaces

import (
	"os"

	"github.com/rs/zerolog/log"
)

// Бизнес-логика, частью которой является интерфейс.
type uploader interface {
	upload(b []byte) (uri string, err error)
}

func uploadFile(u uploader, fileName string) (uri string, err error) {
	log.Info().Msgf("uploader type is: %T", u)
	b, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	return u.upload(b)
}

// Различные варианты имплементации интерфейса.
// На практике, каждый вариант будет в отдельном пакете.
type s3uploader struct{}

// Пример "утиной типизации": тип данных `s3uploader` выполняет контракт
// интерфейса `uploader` неявно, но фактически.
func (s3 *s3uploader) upload(b []byte) (uri string, err error) {
	return
}

type azureuploader struct{}

func (a *azureuploader) upload(b []byte) (uri string, err error) {
	return
}
