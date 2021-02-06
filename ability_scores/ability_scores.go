//go:generate go run ../gen/api_gen.go ability_scores AbilityScore "\"cha\""

package ability_scores

import (
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	AbilityScoreDetail struct {
		Index       string    `json:"index"`
		Name        string    `json:"name"`
		Url         string    `json:"url"`
		FullName    string    `json:"full_name"`
		Description []string  `json:"desc"`
		Skills      []api.Ref `json:"skills"`
	}

	AbilityBonus struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
		Bonus int    `json:"bonus"`
	}
)
