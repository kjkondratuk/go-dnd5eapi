package features

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for features
	Endpoint      = "/features"
	ChildEndpoint = Endpoint + "/"
)

type (
	FeatureDetail struct {
		Index       string   `json:"index"`
		Name        string   `json:"name"`
		Url         string   `json:"url"`
		Class       api.Ref  `json:"class"`
		Level       int      `json:"level"`
		Description []string `json:"desc"`
	}

	featureClient struct {
		basicsProvider api.BasicsProvider
	}

	FeatureClient interface {
		GetList() (*api.ListResponse, error)
		GetByIndex(index string) (*FeatureDetail, error)
	}
)

func NewClient(basicsProvider api.BasicsProvider) FeatureClient {
	return &featureClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *featureClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *featureClient) GetByIndex(index string) (*FeatureDetail, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := FeatureDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
