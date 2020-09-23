package go_dnd5eapi

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/constants"
)

type (
	LanguageChoice struct {
		Choose int      `json:"choose"`
		Type   string   `json:"type"`
		From   []APIRef `json:"from"`
	}

	LanguageDetail struct {
		Index           string   `json:"index"`
		Name            string   `json:"name"`
		Url             string   `json:"url"`
		Type            string   `json:"type"`
		TypicalSpeakers []string `json:"typical_speakers"`
		Script          string   `json:"script"`
	}
)

func (ac *apiClient) GetLanguageList() (*ListResponse, error) {
	return ac.getListForUrl(constants.LanguagesEndpoint)
}

func (ac *apiClient) GetLanguageByIndex(index string) (*LanguageDetail, error) {
	result, err := ac.apiGet(constants.LanguagesChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := LanguageDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
