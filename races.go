package go_dnd5eapi

import (
	"github.com/kjkondratuk/go-dnd5eapi/constants"
	"github.com/kjkondratuk/go-dnd5eapi/response"
)

func (ac *apiClient) GetRaceList() (*response.ListResponse, error) {
	return ac.getListForUrl(constants.RacesEndpoint)
}

func (ac *apiClient) GetRaceByIndex(index string) (*response.RaceDetail, error) {

}
