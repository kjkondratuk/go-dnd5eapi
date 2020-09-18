package go_dnd5eapi

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/constants"
)

type (
	AbilityScoreDetail struct {
		Index       string   `json:"index"`
		Name        string   `json:"name"`
		Url         string   `json:"url"`
		FullName    string   `json:"full_name"`
		Description []string `json:"desc"`
		Skills      []APIRef `json:"skills"`
	}

	AbilityBonus struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
		Bonus int    `json:"bonus"`
	}
)


func (ac *apiClient) GetAbilityScoreList() (*ListResponse, error) {
	return ac.getListForUrl(constants.AbilityScoreEndpoint)
}

func (ac *apiClient) GetAbilityScoreByIndex(index string) (*AbilityScoreDetail, error) {
	result, err := ac.apiGet(constants.AbilityScoreChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := AbilityScoreDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
