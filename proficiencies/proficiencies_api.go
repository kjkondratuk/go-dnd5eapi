// This file is generated by go-dnd5eapi/gen/api_gen.go.  DO NOT EDIT THIS FILE.  Generation parameters are:
// Package Name: proficiencies
// Endpoint: proficiencies
// API Name: Proficiency
// LC API Name: proficiency
// API Class: ProficiencyDetail

package proficiencies

import (
	json "github.com/json-iterator/go"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for ability scores
	Endpoint      = "/proficiencies"
	ChildEndpoint = Endpoint + "/"
)

type (
	proficiencyClient struct {
		basicsProvider api.BasicsProvider
	}
)

type ProficiencyClient interface {
	GetList() (*api.ListResponse, error)
	GetByIndex(index string) (*ProficiencyDetail, error)
}

func NewClient(basicsProvider api.BasicsProvider) ProficiencyClient {
	return &proficiencyClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *proficiencyClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *proficiencyClient) GetByIndex(index string) (*ProficiencyDetail, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := ProficiencyDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
