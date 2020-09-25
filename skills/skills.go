package skills

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for skills
	Endpoint      = "/skills"
	ChildEndpoint = Endpoint + "/"
)

type (
	SkillDetail struct {
		Index        string   `json:"index"`
		Name         string   `json:"name"`
		Url          string   `json:"url"`
		Description  []string `json:"desc"`
		AbilityScore api.Ref  `json:"ability_score"`
	}

	skillClient struct {
		basicsProvider api.BasicsProvider
	}

	SkillClient interface {
		GetList() (*api.ListResponse, error)
		GetByIndex(index string) (*SkillDetail, error)
	}
)

func NewClient(basicsProvider api.BasicsProvider) SkillClient {
	return &skillClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *skillClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *skillClient) GetByIndex(index string) (*SkillDetail, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := SkillDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
