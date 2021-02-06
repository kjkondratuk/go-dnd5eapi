//go:generate go run ../gen/api_gen.go Race "\"dragonborn\""

package races

import (
	"github.com/kjkondratuk/go-dnd5eapi/ability_scores"
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	RaceDetail struct {
		Index                 string                        `json:"index"`
		Name                  string                        `json:"name"`
		Url                   string                        `json:"url"`
		Speed                 int                           `json:"speed"`
		AbilityBonuses        []ability_scores.AbilityBonus `json:"ability_bonuses"`
		Alignment             string                        `json:"alignment"`
		Age                   string                        `json:"age"`
		Size                  string                        `json:"size"`
		SizeDescription       string                        `json:"size_description"`
		StartingProficiencies []api.Ref                     `json:"starting_proficiencies"`
		Languages             []api.Ref                     `json:"languages"`
		LanguageDescription   string                        `json:"language_desc"`
		Traits                []api.Ref                     `json:"traits"`
		TraitOptions          TraitChoice                   `json:"trait_options"`
		Subraces              []api.Ref                     `json:"subraces"`
	}

	TraitChoice struct {
		Choose int       `json:"choose"`
		Type   string    `json:"type"`
		From   []api.Ref `json:"from"`
	}
)
