package go_dnd5eapi

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/constants"
)

func (ac *apiClient) GetEndpointList() (*EndpointResponse, error) {
	result, err := ac.apiGet(constants.RootEndpoint)
	if err != nil {
		return nil, err
	}

	d := EndpointResponse{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
