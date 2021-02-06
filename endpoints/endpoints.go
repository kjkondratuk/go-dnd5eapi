package endpoints

import (
	json "github.com/json-iterator/go"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	EndpointResponse map[string]string

	endpointsClient struct {
		basicsProvider api.BasicsProvider
	}

	EndpointsClient interface {
		GetList() (*EndpointResponse, error)
	}
)

func NewClient(basicsProvider api.BasicsProvider) EndpointsClient {
	return &endpointsClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *endpointsClient) GetList() (*EndpointResponse, error) {
	result, err := ac.basicsProvider.ApiGet("/")
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
