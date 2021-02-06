// This file is generated by go-dnd5eapi/gen/api_gen.go.  DO NOT EDIT THIS FILE.  Generation parameters are:
// Package Name: damage_types
// Endpoint: damage-types
// API Name: DamageType
// LC API Name: damageType
// API Class: api.Description

package damage_types

import (
	json "github.com/json-iterator/go"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for ability scores
	Endpoint      = "/damage-types"
	ChildEndpoint = Endpoint + "/"
)

type (
	damageTypeClient struct {
		basicsProvider api.BasicsProvider
	}
)

type DamageTypeClient interface {
	GetList() (*api.ListResponse, error)
	GetByIndex(index string) (*api.Description, error)
}

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
