package go_dnd5eapi

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/constants"
)

type (
	SkillDetail struct {
		Index        string   `json:"index"`
		Name         string   `json:"name"`
		Url          string   `json:"url"`
		Description  []string `json:"desc"`
		AbilityScore APIRef   `json:"ability_score"`
	}
)

func (ac *apiClient) GetSkillList() (*ListResponse, error) {
	return ac.getListForUrl(constants.SkillsEndpoint)
}

func (ac *apiClient) GetSkillByIndex(index string) (*SkillDetail, error) {
	result, err := ac.apiGet(constants.SkillsChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := SkillDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
