package options

import (
	_ "embed"
	"encoding/json"
)

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Usage string `json:"usage,omitempty"`
	Note  string `json:"nb,omitempty"`
}

//go:embed assets/primary-options.json
var primaryJSON []byte

func (o Option) HasUsage() bool {
	return o.Usage != ""
}

type List []Option

func NewList() *List {
	var options List
	json.Unmarshal(primaryJSON, &options)
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
