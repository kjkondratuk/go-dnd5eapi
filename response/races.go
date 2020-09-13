package response

type (
	RaceDetail struct {
		Index          string `json:"index"`
		Name           string `json:"name"`
		Url            string `json:"url"`
		Speed          int    `json:"speed"`
		AbilityBonuses []AbilityBonus
		Alignment      string `json:"alignment"`
		Age            string `json:"age"`
	}
)
