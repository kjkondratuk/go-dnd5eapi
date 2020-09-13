package response

type (
	Trait struct {
		Index    string   `json:"index"`
		Name     string   `json:"name"`
		Url      string   `json:"url"`
		Races    []APIRef `json:"races"`
		Subraces []APIRef `json:"subraces"`
	}

	TraitChoice struct {
		Choose int      `json:"choose"`
		Type   string   `json:"type"`
		From   []APIRef `json:"from"`
	}
)
