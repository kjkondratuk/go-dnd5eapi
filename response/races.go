package response

type (
	RaceDetail struct {
		Index                 string         `json:"index"`
		Name                  string         `json:"name"`
		Url                   string         `json:"url"`
		Speed                 int            `json:"speed"`
		AbilityBonuses        []AbilityBonus `json:"ability_bonuses"`
		Alignment             string         `json:"alignment"`
		Age                   string         `json:"age"`
		Size                  string         `json:"size"`
		SizeDescription       string         `json:"size_description"`
		StartingProficiencies []string       `json:"starting_proficiencies"`
		Languages             []APIRef       `json:"languages"`
		LanguageDescription   string         `json:"language_desc"`
		Traits                []APIRef       `json:"traits"`
		TraitOptions          []TraitChoice  `json:"trait_options"`
		Subraces              []APIRef       `json:"subraces"`
	}
)
