package ability_scores

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for ability scores
	AbilityScoreEndpoint      = "/ability-scores"
	AbilityScoreChildEndpoint = AbilityScoreEndpoint + "/"
)

type (
	AbilityScoreDetail struct {
		Index       string    `json:"index"`
		Name        string    `json:"name"`
		Url         string    `json:"url"`
		FullName    string    `json:"full_name"`
		Description []string  `json:"desc"`
		Skills      []api.Ref `json:"skills"`
	}

	AbilityBonus struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
		Bonus int    `json:"bonus"`
	}

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
	return ac.basicsProvider.GetListForUrl(AbilityScoreEndpoint)
}

func (ac *abilityScoreClient) GetByIndex(index string) (*AbilityScoreDetail, error) {
	result, err := ac.basicsProvider.ApiGet(AbilityScoreChildEndpoint + index)
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
