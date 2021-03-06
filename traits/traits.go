//go:generate go run ../gen/api_gen.go Trait "\"brave\""

package traits

import "github.com/kjkondratuk/go-dnd5eapi/api"

type (
	TraitDetail struct {
		Index    string    `json:"index"`
		Name     string    `json:"name"`
		Url      string    `json:"url"`
		Races    []api.Ref `json:"races"`
		Subraces []api.Ref `json:"subraces"`
	}
)
