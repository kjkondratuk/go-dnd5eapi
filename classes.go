package go_dnd5eapi

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/constants"
	"github.com/kjkondratuk/go-dnd5eapi/response"
)

func (ac *apiClient) GetClassList() (*response.ListResponse, error) {
	return ac.getListForUrl(constants.ClassEndpoint)
}

func (ac *apiClient) GetClassByIndex(index string) (*response.ClassDetail, error) {
	result, err := ac.apiGet(constants.ClassChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := response.ClassDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
