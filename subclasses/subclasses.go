package subclasses

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for subclasses
	SubclassesEndpoint     = "/subclasses"
	SubclassesChildEnpoint = SubclassesEndpoint + "/"
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
		GetSubclassList() (*api.ListResponse, error)
		GetSubclassByIndex(index string) (*SubclassDetail, error)
	}
)

func NewSubclassClient(basicsProvider api.BasicsProvider) SubclassClient {
	return &subclassClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *subclassClient) GetSubclassList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(SubclassesEndpoint)
}

func (ac *subclassClient) GetSubclassByIndex(index string) (*SubclassDetail, error) {
	result, err := ac.basicsProvider.ApiGet(SubclassesChildEnpoint + index)
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
