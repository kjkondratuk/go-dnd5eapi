// This file is generated by go-dnd5eapi/gen/api_gen.go.  DO NOT EDIT THIS FILE.  Generation parameters are:
// Package Name: traits
// Endpoint: traits
// API Name: Trait
// LC API Name: trait
// API Class: TraitDetail

package traits

import (
	json "github.com/json-iterator/go"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	traitClient struct {
		api.BasicsProvider
		uri string
	}
)

type TraitClient interface {
	GetList() (*api.ListResponse, error)
	QueryList(query map[string]string) (*api.ListResponse, error)
	GetByIndex(index string) (*TraitDetail, error)
}

func NewClient(basicsProvider api.BasicsProvider) TraitClient {
	return &traitClient{
		BasicsProvider: basicsProvider,
		uri: "/traits",
	}
}

func (ac *traitClient) GetList() (*api.ListResponse, error) {
	return ac.GetListForUrl(ac.uri)
}

func (ac *traitClient) QueryList(query map[string]string) (*api.ListResponse, error) {
	return ac.QueryListForUrl(ac.uri, query)
}

func (ac *traitClient) GetByIndex(index string) (*TraitDetail, error) {
	result, err := ac.ApiGet(ac.uri + "/" + index)
	if err != nil {
		return nil, err
	}

	d := TraitDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
