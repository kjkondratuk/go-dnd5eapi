package go_dnd5eapi

import (
	"github.com/kjkondratuk/go-dnd5eapi/constants"
)

type (
	SubclassDetail struct {
		Index       string   `json:"index"`
		Name        string   `json:"name"`
		Url         string   `json:"url"`
		FlavorText  string   `json:"subclass_flavor"`
		Description []string `json:"desc"`
		Class       APIRef   `json:"class"`
		LevelsLink  string   `json:"subclass_levels"`
	}
)

func (ac *apiClient) GetSubclassList() (*ListResponse, error) {
	return ac.getListForUrl(constants.SubclassesEndpoint)
}
