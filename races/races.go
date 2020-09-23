package races

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/ability_scores"
	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/kjkondratuk/go-dnd5eapi/constants"
)

type (
	RaceDetail struct {
		Index                 string                        `json:"index"`
		Name                  string                        `json:"name"`
		Url                   string                        `json:"url"`
		Speed                 int                           `json:"speed"`
		AbilityBonuses        []ability_scores.AbilityBonus `json:"ability_bonuses"`
		Alignment             string                        `json:"alignment"`
		Age                   string                        `json:"age"`
		Size                  string                        `json:"size"`
		SizeDescription       string                        `json:"size_description"`
		StartingProficiencies []api.APIRef                  `json:"starting_proficiencies"`
		Languages             []api.APIRef                  `json:"languages"`
		LanguageDescription   string                        `json:"language_desc"`
		Traits                []api.APIRef                  `json:"traits"`
		TraitOptions          []TraitChoice                 `json:"trait_options"`
		Subraces              []api.APIRef                  `json:"subraces"`
	}

	TraitChoice struct {
		Choose int          `json:"choose"`
		Type   string       `json:"type"`
		From   []api.APIRef `json:"from"`
	}

	raceClient struct {
		basicsProvider api.BasicsProvider
	}

	RaceClient interface {
		GetRaceList() (*api.ListResponse, error)
		GetRaceByIndex(index string) (*RaceDetail, error)
	}
)

func NewRaceClient(basicsProvider api.BasicsProvider) RaceClient {
	return &raceClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *raceClient) GetRaceList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(constants.RacesEndpoint)
}

func (ac *raceClient) GetRaceByIndex(index string) (*RaceDetail, error) {
	result, err := ac.basicsProvider.ApiGet(constants.RacesChildEndpoint + index)
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
