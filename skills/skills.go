package skills

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for skills
	SkillsEndpoint      = "/skills"
	SkillsChildEndpoint = SkillsEndpoint + "/"
)

type (
	SkillDetail struct {
		Index        string     `json:"index"`
		Name         string     `json:"name"`
		Url          string     `json:"url"`
		Description  []string   `json:"desc"`
		AbilityScore api.APIRef `json:"ability_score"`
	}

	skillClient struct {
		basicsProvider api.BasicsProvider
	}

	SkillClient interface {
		GetSkillList() (*api.ListResponse, error)
		GetSkillByIndex(index string) (*SkillDetail, error)
	}
)

func NewSkillClient(basicsProvider api.BasicsProvider) SkillClient {
	return &skillClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *skillClient) GetSkillList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(SkillsEndpoint)
}

func (ac *skillClient) GetSkillByIndex(index string) (*SkillDetail, error) {
	result, err := ac.basicsProvider.ApiGet(SkillsChildEndpoint + index)
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
