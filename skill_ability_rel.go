package go_dnd5eapi

import (
	"errors"
)

func (ac *apiClient) GetSkillsForAbilityByIndex(index string) ([]SkillDetail, error) {
	ability, err := ac.GetAbilityScoreByIndex(index)
	if err != nil {
		return nil, err
	}

	var result []SkillDetail
	for _, v := range ability.Skills {
		skill, err := ac.GetSkillByIndex(v.Index)
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

func (ac *apiClient) GetAbilityScoreForSkillByIndex(skillIndex string) (*AbilityScoreDetail, error) {
	skill, err := ac.GetSkillByIndex(skillIndex)
	if err != nil {
		return nil, err
	}
	abilityScore, err := ac.GetAbilityScoreByIndex(skill.AbilityScore.Index)
	if err != nil {
		return nil, err
	}
	return abilityScore, nil
}
