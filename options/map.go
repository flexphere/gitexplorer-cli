package options

import (
	"encoding/json"
	"io/ioutil"
)

type Map map[string]List

func NewMap(file string) *Map {
	var options Map
	byteValue, _ := ioutil.ReadFile(file)
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
