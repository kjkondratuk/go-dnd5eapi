//go:generate go run ../gen/api_gen.go ../gen features Feature FeatureDetail "\"arcane-tradition\""

package features

import (
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	FeatureDetail struct {
		Index       string   `json:"index"`
		Name        string   `json:"name"`
		Url         string   `json:"url"`
		Class       api.Ref  `json:"class"`
		Level       int      `json:"level"`
		Description []string `json:"desc"`
	}
)
