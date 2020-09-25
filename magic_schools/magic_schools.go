package magic_schools

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for magic schools
	Endpoint      = "/magic-schools"
	ChildEndpoint = Endpoint + "/"
)

type (
	MagicSchoolDetail struct {
		Index       string  `json:"index"`
		Name        string  `json:"name"`
		Url         string  `json:"url"`
		Class       api.Ref `json:"class"`
		Level       int     `json:"level"`
		Description string  `json:"desc"`
	}

	magicSchoolsClient struct {
		basicsProvider api.BasicsProvider
	}

	MagicSchoolsClient interface {
		GetList() (*api.ListResponse, error)
		GetByIndex(index string) (*MagicSchoolDetail, error)
	}
)

func NewClient(basicsProvider api.BasicsProvider) MagicSchoolsClient {
	return &magicSchoolsClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *magicSchoolsClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *magicSchoolsClient) GetByIndex(index string) (*MagicSchoolDetail, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := MagicSchoolDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
