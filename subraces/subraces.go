//go:generate go run ../gen/api_gen.go Subrace "\"high-elf\""

package subraces

import (
	"github.com/kjkondratuk/go-dnd5eapi/ability_scores"
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	SubraceDetail struct {
		Index                 string                        `json:"index"`
		Name                  string                        `json:"name"`
		Url                   string                        `json:"url"`
		Race                  api.Ref                       `json:"race"`
		Description           string                        `json:"desc"`
		AbilityBonuses        []ability_scores.AbilityBonus `json:"ability_bonuses"`
		StartingProficiencies []api.Ref                     `json:"starting_proficiencies"`
		Languages             []api.Ref                     `json:"languages"`
		LanguageOptions       api.Choice                    `json:"language_options"`
		RacialTraits          []api.Ref                     `json:"racial_traits"`
		RacialTraitOptions    api.Choice                    `json:"racial_trait_options"`
	}
)
