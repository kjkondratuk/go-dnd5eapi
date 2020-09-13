package response

type (
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
