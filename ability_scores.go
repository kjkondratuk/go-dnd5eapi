package go_dnd5eapi

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/constants"
	"github.com/kjkondratuk/go-dnd5eapi/response"
)

func (ac *apiClient) GetAbilityScoreList() (*response.ListResponse, error) {
	return ac.getListForUrl(constants.AbilityScoreEndpoint)
}

func (ac *apiClient) GetAbilityScoreByIndex(index string) (*response.AbilityScoreDetail, error) {
	result, err := ac.apiGet(constants.AbilityScoreChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := response.AbilityScoreDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
