package proficiencies

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for proficiencies
	ProficienciesEndpoint      = "/proficiencies"
	ProficienciesChildEndpoint = ProficienciesEndpoint + "/"
)

type (
	ProficiencyDetail struct {
		Index      string       `json:"index"`
		Name       string       `json:"name"`
		Url        string       `json:"url"`
		Type       string       `json:"type"`
		Classes    []api.APIRef `json:"classes"`
		Races      []api.APIRef `json:"races"`
		References []api.APIRef `json:"references"`
	}

	proficiencyClient struct {
		basicsProvider api.BasicsProvider
	}

	ProficiencyClient interface {
		GetProficiencyList() (*api.ListResponse, error)
		GetProficiencyByIndex(index string) (*ProficiencyDetail, error)
	}
)

func NewProficiencyClient(basicsProvider api.BasicsProvider) ProficiencyClient {
	return &proficiencyClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *proficiencyClient) GetProficiencyList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(ProficienciesEndpoint)
}

func (ac *proficiencyClient) GetProficiencyByIndex(index string) (*ProficiencyDetail, error) {
	result, err := ac.basicsProvider.ApiGet(ProficienciesChildEndpoint + index)
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
