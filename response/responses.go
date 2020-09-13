package response

type (
	EndpointResponse map[string]string

	APIRef struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	}

	ListResponse struct {
		Count   int `json:"count"`
		Results []APIRef
	}
)
