package response

type (
	ClassDetail struct {
		Index                 string              `json:"index"`
		Name                  string              `json:"name"`
		Url                   string              `json:"url"`
		HitDie                int                 `json:"hit_die"`
		ProficiencyChoices    []ProficiencyChoice `json:"proficiency_choices"`
		Proficiencies         []APIRef            `json:"proficiencies"`
		SavingThrows          APIRef              `json:"saving_throws"`
		StartingEquipmentLink string              `json:"starting_equipment"`
		LevelsLink            string              `json:"class_levels"`
		Subclasses            []APIRef            `json:"subclasses"`
	}
)
