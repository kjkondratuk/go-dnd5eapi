package subclasses

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for subclasses
	Endpoint     = "/subclasses"
	ChildEnpoint = Endpoint + "/"
)

type (
	SubclassDetail struct {
		Index       string     `json:"index"`
		Name        string     `json:"name"`
		Url         string     `json:"url"`
		FlavorText  string     `json:"subclass_flavor"`
		Description []string   `json:"desc"`
		Class       api.APIRef `json:"class"`
		LevelsLink  string     `json:"subclass_levels"`
	}

	subclassClient struct {
		basicsProvider api.BasicsProvider
	}

	SubclassClient interface {
		GetList() (*api.ListResponse, error)
		GetByIndex(index string) (*SubclassDetail, error)
	}
)

func NewClient(basicsProvider api.BasicsProvider) SubclassClient {
	return &subclassClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *subclassClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *subclassClient) GetByIndex(index string) (*SubclassDetail, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEnpoint + index)
	if err != nil {
		return nil, err
	}

	d := SubclassDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
