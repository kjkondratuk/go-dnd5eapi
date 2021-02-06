// This file is generated by go-dnd5eapi/gen/api_gen.go.  DO NOT EDIT THIS FILE.  Generation parameters are:
// Package Name: ability_scores
// Endpoint: ability-scores
// API Name: AbilityScore
// LC API Name: abilityScore
// API Class: AbilityScoreDetail

package ability_scores

import (
	json "github.com/json-iterator/go"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	abilityScoreClient struct {
		api.BasicsProvider
		uri string
	}
)

type AbilityScoreClient interface {
	GetList() (*api.ListResponse, error)
	QueryList(query map[string]string) (*api.ListResponse, error)
	GetByIndex(index string) (*AbilityScoreDetail, error)
}

func NewClient(basicsProvider api.BasicsProvider) AbilityScoreClient {
	return &abilityScoreClient{
		BasicsProvider: basicsProvider,
		uri: "/ability-scores",
	}
}

func (ac *abilityScoreClient) GetList() (*api.ListResponse, error) {
	return ac.GetListForUrl(ac.uri)
}

func (ac *abilityScoreClient) QueryList(query map[string]string) (*api.ListResponse, error) {
	return ac.QueryListForUrl(ac.uri, query)
}

func (ac *abilityScoreClient) GetByIndex(index string) (*AbilityScoreDetail, error) {
	result, err := ac.ApiGet(ac.uri + "/" + index)
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
