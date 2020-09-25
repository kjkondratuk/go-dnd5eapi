package monsters

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for monsters
	Endpoint      = "/monsters"
	ChildEndpoint = Endpoint + "/"
)

type (
	MonsterDetail struct {
		Index                 string             `json:"index"`
		Name                  string             `json:"name"`
		Url                   string             `json:"url"`
		Size                  string             `json:"size"`
		Type                  string             `json:"type"`
		Subtype               string             `json:"subtype"`
		Alignment             string             `json:"alignment"`
		ArmorClass            int                `json:"armor_class"`
		HitPoints             int                `json:"hit_points"`
		HitDice               string             `json:"hit_dice"`
		Speed                 Speed              `json:"speed"`
		Strength              int                `json:"strength"`
		Dexterity             int                `json:"dexterity"`
		Constitution          int                `json:"constitution"`
		Intelligence          int                `json:"intelligence"`
		Wisdom                int                `json:"wisdom"`
		Charisma              int                `json:"charisma"`
		Proficiencies         []ProficiencyValue `json:"proficiencies"`
		DamageVulnerabilities []string           `json:"damage_vulnerabilities"`
		DamageResistances     []string           `json:"damage_resistances"`
		DamageImmunities      []string           `json:"damage_immunities"`
		ConditionImmunities   []string           `json:"condition_immunities"`
		Senses                Sense              `json:"senses"`
		Languages             string             `json:"languages"`
		ChallengeRating       int                `json:"challenge_rating"`
		SpecialAbilities      []SpecialAbility   `json:"special_abilities"`
	}

	SpecialAbility struct {
		Name        string          `json:"name"`
		Description api.Description `json:"desc"`
		Usage       Usage           `json:"usage"`
	}

	Usage struct {
		Type  string `json:"type"`
		Times int    `json:"times"`
	}

	Sense struct {
		Blindsight        string `json:"blindsight"`
		Darkvision        string `json:"darkvision"`
		PassivePerception string `json:"passive_perception"`
	}

	ProficiencyValue struct {
		Value       int     `json:"value"`
		Proficiency api.Ref `json:"proficiency"`
	}

	Speed struct {
		Walk  string `json:"walk"`
		Climb string `json:"climb"`
		Fly   string `json:"fly"`
	}

	monsterClient struct {
		basicsProvider api.BasicsProvider
	}

	MonsterClient interface {
		GetList() (*api.ListResponse, error)
		GetByIndex(index string) (*MonsterDetail, error)
	}
)

func NewClient(basicsProvider api.BasicsProvider) MonsterClient {
	return &monsterClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *monsterClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *monsterClient) GetByIndex(index string) (*MonsterDetail, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := MonsterDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
