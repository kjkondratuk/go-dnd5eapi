package response

type (
	TraitChoice struct {
		Choose int      `json:"choose"`
		Type   string   `json:"type"`
		From   []APIRef `json:"from"`
	}
)
