package response

type (
	EndpointResponse map[string]string

	APIRef struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	}

	AbilityScoreDetail struct {
		Index       string   `json:"index"`
		Name        string   `json:"name"`
		Url         string   `json:"url"`
		FullName    string   `json:"full_name"`
		Description []string `json:"desc"`
		Skills      []APIRef `json:"skills"`
	}

	// endregion

	// region Skills
	SkillDetail struct {
		Index        string   `json:"index"`
		Name         string   `json:"name"`
		Url          string   `json:"url"`
		Description  []string `json:"desc"`
		AbilityScore APIRef   `json:"ability_score"`
	}

	ListResponse struct {
		Count   int `json:"count"`
		Results []APIRef
	}

	// endregion

	// region Class
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

	// endregion

	// region Proficiency
	ProficiencyChoice struct {
		Choose int      `json:"choose"`
		Type   string   `json:"type"`
		From   []APIRef `json:"from"`
	}

	ProficiencyDetail struct {
		Index      string   `json:"index"`
		Name       string   `json:"name"`
		Url        string   `json:"url"`
		Type       string   `json:"type"`
		Classes    []APIRef `json:"classes"`
		Races      []APIRef `json:"races"`
		References []APIRef `json:"references"`
	}
	// endregion

	// region Subclass
	SubclassDetail struct {
		Index       string   `json:"index"`
		Name        string   `json:"name"`
		Url         string   `json:"url"`
		FlavorText  string   `json:"subclass_flavor"`
		Description []string `json:"desc"`
		Class       APIRef   `json:"class"`
		LevelsLink  string   `json:"subclass_levels"`
	}
	// endregion
)
