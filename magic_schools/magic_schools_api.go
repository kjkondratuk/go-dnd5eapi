// This file is generated by go-dnd5eapi/gen/api_gen.go.  DO NOT EDIT THIS FILE.  Generation parameters are:
// Package Name: magic_schools
// Endpoint: magic-schools
// API Name: MagicSchool
// LC API Name: magicSchool
// API Class: MagicSchoolDetail

package magic_schools

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for ability scores
	Endpoint      = "/magic-schools"
	ChildEndpoint = Endpoint + "/"
)

type (
	magicSchoolClient struct {
		basicsProvider api.BasicsProvider
	}
)

type MagicSchoolClient interface {
	GetList() (*api.ListResponse, error)
	GetByIndex(index string) (*MagicSchoolDetail, error)
}

func NewClient(basicsProvider api.BasicsProvider) MagicSchoolClient {
	return &magicSchoolClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *magicSchoolClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *magicSchoolClient) GetByIndex(index string) (*MagicSchoolDetail, error) {
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