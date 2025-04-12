package interfaces

import (
	"encoding/json"
	"strconv"

	"github.com/rs/zerolog/log"
)

func anyType() {
	jsonb := []byte(`{"key1": 1, "key2": "2"}`)

	var values map[string]any

	json.Unmarshal(jsonb, &values)

	for k, v := range values {
		var value int
		switch val := v.(type) {
		case float64:
			value = int(val)
		case string:
			value, _ = strconv.Atoi(val)
		}
		log.Info().Msgf("key=%v\tvalue=%v", k, value)
	}
}

func anyParam(params ...any) {
	for _, param := range params {
		log.Info().Msgf("param=%v type=%T", param, param)
	}
}
