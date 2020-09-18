package go_dnd5eapi

type (
	Subrace struct {
		Index                 string           `json:"index"`
		Name                  string           `json:"name"`
		Url                   string           `json:"url"`
		AbilityBonuses        []AbilityBonus   `json:"ability_bonuses"`
		StartingProficiencies []APIRef         `json:"starting_proficiencies"`
		Languages             []APIRef         `json:"languages"`
		LanguageChoice        []LanguageChoice `json:"language_options"`
		RacialTraits          []Trait          `json:"racial_traits"`
	}
)
