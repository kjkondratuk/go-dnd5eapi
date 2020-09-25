package categories

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for equipment categories
	Endpoint      = "/equipment-categories"
	ChildEndpoint = Endpoint + "/"
)

type (
	CategoryDetail struct {
		Index     string    `json:"index"`
		Name      string    `json:"name"`
		Url       string    `json:"url"`
		Equipment []api.Ref `json:"equipment"`
	}

	categoryClient struct {
		basicsProvider api.BasicsProvider
	}

	CategoryClient interface {
		GetList() (*api.ListResponse, error)
		GetByIndex(index string) (*CategoryDetail, error)
	}
)

func NewClient(basicsProvider api.BasicsProvider) CategoryClient {
	return &categoryClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *categoryClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *categoryClient) GetByIndex(index string) (*CategoryDetail, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := CategoryDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
