//go:generate go run ../gen/api_gen.go proficiencies Proficiency "\"battleaxes\""

package proficiencies

import (
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	ProficiencyDetail struct {
		Index      string    `json:"index"`
		Name       string    `json:"name"`
		Url        string    `json:"url"`
		Type       string    `json:"type"`
		Classes    []api.Ref `json:"classes"`
		Races      []api.Ref `json:"races"`
		References []api.Ref `json:"references"`
	}
)
