package go_dnd5eapi

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/constants"
)

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

func (ac *apiClient) GetProficiencyList() (*ListResponse, error) {
	return ac.getListForUrl(constants.ProficienciesEndpoint)
}

func (ac *apiClient) GetProficiencyByIndex(index string) (*ProficiencyDetail, error) {
	result, err := ac.apiGet(constants.ProficienciesChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := ProficiencyDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
