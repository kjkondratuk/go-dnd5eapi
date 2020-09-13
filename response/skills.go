package response

type (
	SkillDetail struct {
		Index        string   `json:"index"`
		Name         string   `json:"name"`
		Url          string   `json:"url"`
		Description  []string `json:"desc"`
		AbilityScore APIRef   `json:"ability_score"`
	}
)
