package response

type (
	EndpointResponse map[string]string

	APIRef struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	}

	AbilityScoreDetail struct {
		Index       string         `json:"index"`
		Name        string         `json:"name"`
		Url         string         `json:"url"`
		FullName    string         `json:"full_name"`
		Description []string       `json:"desc"`
		Skills      []APIRef `json:"skills"`
	}

	// endregion

	// region Skills
	SkillDetail struct {
		Index        string `json:"index"`
		Name         string `json:"name"`
		Url          string `json:"url"`
		Description  []string `json:"desc"`
		AbilityScore APIRef `json:"ability_score"`
	}

	ListResponse struct {
		Count int `json:"count"`
		Results []APIRef
	}

	// endregion

	ClassDetail struct {
		Index              string              `json:"index"`
		Name               string              `json:"name"`
		Url                string              `json:"url"`
		HitDie             int                 `json:"hit_die"`
		ProficiencyChoices []ProficiencyChoice `json:"proficiency_choices"`
	}

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
)
