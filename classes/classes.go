package classes

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for classes
	Endpoint      = "/classes"
	ChildEndpoint = Endpoint + "/"
)

type (
	ClassDetail struct {
		Index                 string          `json:"index"`
		Name                  string          `json:"name"`
		Url                   string          `json:"url"`
		HitDie                int             `json:"hit_die"`
		ProficiencyChoices    []api.APIChoice `json:"proficiency_choices"`
		Proficiencies         []api.APIRef    `json:"proficiencies"`
		SavingThrows          []api.APIRef    `json:"saving_throws"`
		StartingEquipmentLink string          `json:"starting_equipment"`
		LevelsLink            string          `json:"class_levels"`
		Subclasses            []api.APIRef    `json:"subclasses"`
	}

	classesClient struct {
		basicsProvider api.BasicsProvider
	}

	ClassesClient interface {
		GetList() (*api.ListResponse, error)
		GetByIndex(index string) (*ClassDetail, error)
	}
)

func NewClient(basicsProvider api.BasicsProvider) ClassesClient {
	return &classesClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *classesClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *classesClient) GetByIndex(index string) (*ClassDetail, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := ClassDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
