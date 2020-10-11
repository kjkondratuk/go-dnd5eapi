//go:generate go run ../gen/api_gen.go ../gen languages Language LanguageDetail "\"abyssal\""

package languages

import (
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	LanguageChoice struct {
		Choose int       `json:"choose"`
		Type   string    `json:"type"`
		From   []api.Ref `json:"from"`
	}

	LanguageDetail struct {
		Index           string   `json:"index"`
		Name            string   `json:"name"`
		Url             string   `json:"url"`
		Type            string   `json:"type"`
		TypicalSpeakers []string `json:"typical_speakers"`
		Script          string   `json:"script"`
	}
)
