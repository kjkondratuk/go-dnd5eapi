package go_dnd5eapi

import (
	"encoding/json"
	"github.com/kjkondratuk/go-dnd5eapi/constants"
	"github.com/kjkondratuk/go-dnd5eapi/response"
)

func (ac *apiClient) GetRaceList() (*response.ListResponse, error) {
	return ac.getListForUrl(constants.RacesEndpoint)
}

func (ac *apiClient) GetRaceByIndex(index string) (*response.RaceDetail, error) {
	result, err := ac.apiGet(constants.ProficienciesChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := response.RaceDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
