package options

import (
	_ "embed"
	"encoding/json"
)

type Map map[string]List

//go:embed assets/secondary-options.json
var secondaryJSON []byte

//go:embed assets/tertiary-options.json
var tertiaryJSON []byte

func NewMap(key string) *Map {
	var options Map
	var byteValue []byte
	if key == "secondary" {
		byteValue = secondaryJSON
	}
	if key == "tertiary" {
		byteValue = tertiaryJSON
	}
	json.Unmarshal(byteValue, &options)
	return &options
}

func (o Map) GetLabels(key string) []string {
	var result []string
	for _, option := range o[key] {
		result = append(result, option.Label)
	}
	return result
}

func (o Map) GetValue(key string, label string) Option {
	for _, option := range o[key] {
		if option.Label == label {
			return option
		}
	}
	return Option{}
}
