package go_dnd5eapi

import (
	"encoding/json"
	"github.com/kjkondratuk/go-dnd5eapi/constants"
)

type (
	ClassDetail struct {
		Index                 string              `json:"index"`
		Name                  string              `json:"name"`
		Url                   string              `json:"url"`
		HitDie                int                 `json:"hit_die"`
		ProficiencyChoices    []ProficiencyChoice `json:"proficiency_choices"`
		Proficiencies         []APIRef            `json:"proficiencies"`
		SavingThrows          []APIRef              `json:"saving_throws"`
		StartingEquipmentLink string              `json:"starting_equipment"`
		LevelsLink            string              `json:"class_levels"`
		Subclasses            []APIRef            `json:"subclasses"`
	}
)

func (ac *apiClient) GetClassList() (*ListResponse, error) {
	return ac.getListForUrl(constants.ClassEndpoint)
}

func (ac *apiClient) GetClassByIndex(index string) (*ClassDetail, error) {
	result, err := ac.apiGet(constants.ClassChildEndpoint + index)
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

