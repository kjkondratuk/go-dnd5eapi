// This file is generated by go-dnd5eapi/gen/api_gen.go.  DO NOT EDIT THIS FILE.  Generation parameters are:
// Package Name: ability_scores
// Endpoint: ability-scores
// API Name: AbilityScore
// LC API Name: abilityScore
// API Class: AbilityScoreDetail

package ability_scores

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for ability scores
	Endpoint      = "/ability-scores"
	ChildEndpoint = Endpoint + "/"
)

type (
	abilityScoreClient struct {
		basicsProvider api.BasicsProvider
	}
)

type AbilityScoreClient interface {
	GetList() (*api.ListResponse, error)
	GetByIndex(index string) (*AbilityScoreDetail, error)
}

func NewClient(basicsProvider api.BasicsProvider) AbilityScoreClient {
	return &abilityScoreClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *abilityScoreClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *abilityScoreClient) GetByIndex(index string) (*AbilityScoreDetail, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := AbilityScoreDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}