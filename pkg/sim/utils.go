package sim

import (
	"reflect"

	"github.com/go-viper/mapstructure/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func nextId() string {
	return uuid.New().String()
}

func Convert(input any, output any) error {
	// Create a decoder config with a hook to handle float64 to int conversion
	config := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			// Convert float64 to int when the target is int
			func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
				if f.Kind() == reflect.Float64 && t.Kind() == reflect.Int {
					return int(data.(float64)), nil
				}
				return data, nil
			},
		),
		Result: output,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create decoder")
		return err
	}

	err = decoder.Decode(input)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert")
		return err
	}
	return nil
}
