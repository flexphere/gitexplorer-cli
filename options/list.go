package options

import (
	"encoding/json"
	"io/ioutil"
)

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Usage string `json:"usage,omitempty"`
	Note  string `json:"nb,omitempty"`
}

func (o Option) HasUsage() bool {
	return o.Usage != ""
}

type List []Option

func NewList(file string) *List {
	var options List
	byteValue, _ := ioutil.ReadFile(file)
	json.Unmarshal(byteValue, &options)
	return &options
}

func (o List) GetLabels() []string {
	var result []string
	for _, option := range o {
		result = append(result, option.Label)
	}
	return result
}

func (o List) GetValue(label string) string {
	for _, option := range o {
		if option.Label == label {
			return option.Value
		}
	}
	return ""
}
