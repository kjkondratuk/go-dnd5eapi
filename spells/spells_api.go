// This file is generated by go-dnd5eapi/gen/api_gen.go.  DO NOT EDIT THIS FILE.  Generation parameters are:
// Package Name: spells
// Endpoint: spells
// API Name: Spell
// LC API Name: spell
// API Class: SpellDetail

package spells

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for ability scores
	Endpoint      = "/spells"
	ChildEndpoint = Endpoint + "/"
)

type (
	spellClient struct {
		basicsProvider api.BasicsProvider
	}
)

type SpellClient interface {
	GetList() (*api.ListResponse, error)
	GetByIndex(index string) (*SpellDetail, error)
}

func NewClient(basicsProvider api.BasicsProvider) SpellClient {
	return &spellClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *spellClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *spellClient) GetByIndex(index string) (*SpellDetail, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := SpellDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}