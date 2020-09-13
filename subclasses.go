package go_dnd5eapi

import (
	"github.com/kjkondratuk/go-dnd5eapi/constants"
	"github.com/kjkondratuk/go-dnd5eapi/response"
)

func (ac *apiClient) GetSubclassList() (*response.ListResponse, error) {
	return ac.getListForUrl(constants.SubclassesEndpoint)
}
