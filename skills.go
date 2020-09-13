package go_dnd5eapi

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/constants"
	"github.com/kjkondratuk/go-dnd5eapi/response"
)

func (ac *apiClient) GetSkillList() (*response.ListResponse, error) {
	return ac.getListForUrl(constants.SkillsEndpoint)
}

func (ac *apiClient) GetSkillByIndex(index string) (*response.SkillDetail, error) {
	result, err := ac.apiGet(constants.SkillsChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := response.SkillDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
