package response

type (
	EndpointResponse map[string]string

	AllAbilityScoresResponse struct {
		Count   int                   `json:"count"`
		Results []SummaryAbilityScore `json:"results"`
	}

	GenericApiEntity struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	}

	SummaryAbilityScore struct {
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
		Skills      []SummarySkill `json:"skills"`
	}

	SummarySkill struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	}

	SkillDetail struct {
		Index        string `json:"index"`
		Name         string `json:"name"`
		Url          string `json:"url"`
		Description  string `json:"desc"`
		AbilityScore SummaryAbilityScore
	}

	AllClassesResponse struct {
		Count   int
		Results SummaryClass
	}

	SummaryClass struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	}

	ClassSummary struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	}

	ClassDetail struct {
		Index              string              `json:"index"`
		Name               string              `json:"name"`
		Url                string              `json:"url"`
		HitDie             int                 `json:"hit_die"`
		ProficiencyChoices []ProficiencyChoice `json:"proficiency_choices"`
	}

	ProficiencyChoice struct {
		Choose int                  `json:"choose"`
		Type   string               `json:"type"`
		From   []ProficiencySummary `json:"from"`
	}

	ProficiencySummary struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	}

	ProficiencyDetail struct {
		Index      string             `json:"index"`
		Name       string             `json:"name"`
		Url        string             `json:"url"`
		Type       string             `json:"type"`
		Classes    []ClassSummary     `json:"classes"`
		Races      []RaceSummary      `json:"races"`
		References []GenericApiEntity `json:"references"`
	}

	RaceSummary struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	}
)
