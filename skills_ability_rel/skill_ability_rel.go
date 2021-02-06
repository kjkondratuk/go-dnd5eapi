package skills_ability_rel

import (
	"errors"

	"github.com/kjkondratuk/go-dnd5eapi/api/ability_scores"
	"github.com/kjkondratuk/go-dnd5eapi/skills"
)

type (
	skillAbilityRelClient struct {
		sc  skills.SkillClient
		asc ability_scores.AbilityScoreClient
	}

	SkillAbilityRelClient interface {
		GetSkillsForAbilityByIndex(index string) ([]skills.SkillDetail, error)
		GetAbilityScoreForSkillByIndex(abilityScoreIndex string) (*ability_scores.AbilityScoreDetail, error)
	}
)

func NewClient(sc skills.SkillClient, asc ability_scores.AbilityScoreClient) SkillAbilityRelClient {
	return &skillAbilityRelClient{
		sc:  sc,
		asc: asc,
	}
}

func (ac *skillAbilityRelClient) GetSkillsForAbilityByIndex(index string) ([]skills.SkillDetail, error) {
	ability, err := ac.asc.GetByIndex(index)
	if err != nil {
		return nil, err
	}

	var result []skills.SkillDetail
	for _, v := range ability.Skills {
		skill, err := ac.sc.GetByIndex(v.Index)
		if err != nil {
			return nil, err
		}
		if skill != nil {
			result = append(result, *skill)
		} else {
			return nil, errors.New("No skill called " + v.Index + " could be found")
		}
	}
	return result, nil
}

func (ac *skillAbilityRelClient) GetAbilityScoreForSkillByIndex(skillIndex string) (*ability_scores.AbilityScoreDetail, error) {
	skill, err := ac.sc.GetByIndex(skillIndex)
	if err != nil {
		return nil, err
	}
	abilityScore, err := ac.asc.GetByIndex(skill.AbilityScore.Index)
	if err != nil {
		return nil, err
	}
	return abilityScore, nil
}
