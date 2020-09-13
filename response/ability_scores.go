package response

type (
	AbilityScoreDetail struct {
		Index       string   `json:"index"`
		Name        string   `json:"name"`
		Url         string   `json:"url"`
		FullName    string   `json:"full_name"`
		Description []string `json:"desc"`
		Skills      []APIRef `json:"skills"`
	}

	AbilityBonus struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
		Bonus int    `json:"bonus"`
	}
)
