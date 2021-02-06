//go:generate go run ../gen/api_gen.go subclasses Subclass "\"berserker\""

package subclasses

import (
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	SubclassDetail struct {
		Index       string   `json:"index"`
		Name        string   `json:"name"`
		Url         string   `json:"url"`
		FlavorText  string   `json:"subclass_flavor"`
		Description []string `json:"desc"`
		Class       api.Ref  `json:"class"`
		LevelsLink  string   `json:"subclass_levels"`
	}
)
