//go:generate go run ../gen/api_gen.go ../gen classes Class ClassDetail "\"Rogue\""

package classes

import (
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	ClassDetail struct {
		Index                 string       `json:"index"`
		Name                  string       `json:"name"`
		Url                   string       `json:"url"`
		HitDie                int          `json:"hit_die"`
		ProficiencyChoices    []api.Choice `json:"proficiency_choices"`
		Proficiencies         []api.Ref    `json:"proficiencies"`
		SavingThrows          []api.Ref    `json:"saving_throws"`
		StartingEquipmentLink string       `json:"starting_equipment"`
		LevelsLink            string       `json:"class_levels"`
		Subclasses            []api.Ref    `json:"subclasses"`
	}
)
