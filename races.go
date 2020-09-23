package go_dnd5eapi

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/constants"
)

type (
	RaceDetail struct {
		Index                 string         `json:"index"`
		Name                  string         `json:"name"`
		Url                   string         `json:"url"`
		Speed                 int            `json:"speed"`
		AbilityBonuses        []AbilityBonus `json:"ability_bonuses"`
		Alignment             string         `json:"alignment"`
		Age                   string         `json:"age"`
		Size                  string         `json:"size"`
		SizeDescription       string         `json:"size_description"`
		StartingProficiencies []APIRef       `json:"starting_proficiencies"`
		Languages             []APIRef       `json:"languages"`
		LanguageDescription   string         `json:"language_desc"`
		Traits                []APIRef       `json:"traits"`
		TraitOptions          []TraitChoice  `json:"trait_options"`
		Subraces              []APIRef       `json:"subraces"`
	}
)

func (ac *apiClient) GetRaceList() (*ListResponse, error) {
	return ac.getListForUrl(constants.RacesEndpoint)
}

func (ac *apiClient) GetRaceByIndex(index string) (*RaceDetail, error) {
	result, err := ac.apiGet(constants.RacesChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := RaceDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
