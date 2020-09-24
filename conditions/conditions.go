package conditions

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for conditions
	Endpoint      = "/conditions"
	ChildEndpoint = Endpoint + "/"
)

type (
	conditionClient struct {
		basicsProvider api.BasicsProvider
	}

	ConditionClient interface {
		GetList() (*api.ListResponse, error)
		GetByIndex(index string) (*api.APIDescription, error)
	}
)

func NewClient(basicsProvider api.BasicsProvider) ConditionClient {
	return &conditionClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *conditionClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *conditionClient) GetByIndex(index string) (*api.APIDescription, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := api.APIDescription{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
