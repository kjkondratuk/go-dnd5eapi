package damage

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for damage types
	Endpoint      = "/damage-types"
	ChildEndpoint = Endpoint + "/"
)

type (
	damageTypeClient struct {
		basicsProvider api.BasicsProvider
	}

	DamageTypeClient interface {
		GetList() (*api.ListResponse, error)
		GetByIndex(index string) (*api.Description, error)
	}
)

func NewClient(basicsProvider api.BasicsProvider) DamageTypeClient {
	return &damageTypeClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *damageTypeClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *damageTypeClient) GetByIndex(index string) (*api.Description, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := api.Description{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
