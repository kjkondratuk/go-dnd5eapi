package go_dnd5eapi

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/constants"
	"github.com/kjkondratuk/go-dnd5eapi/response"
)

func (ac *apiClient) GetProficiencyByIndex(index string) (*response.ProficiencyDetail, error) {
	result, err := ac.apiGet(constants.ProficienciesChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := response.ProficiencyDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
