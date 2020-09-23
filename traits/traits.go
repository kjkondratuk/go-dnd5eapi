package traits

import "github.com/kjkondratuk/go-dnd5eapi/api"

type (
	Trait struct {
		Index    string       `json:"index"`
		Name     string       `json:"name"`
		Url      string       `json:"url"`
		Races    []api.APIRef `json:"races"`
		Subraces []api.APIRef `json:"subraces"`
	}
)
