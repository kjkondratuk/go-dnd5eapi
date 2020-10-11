//go:generate go run ../gen/api_gen.go ../gen skills Skill SkillDetail "\"acrobatics\""

package skills

import (
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	SkillDetail struct {
		Index        string   `json:"index"`
		Name         string   `json:"name"`
		Url          string   `json:"url"`
		Description  []string `json:"desc"`
		AbilityScore api.Ref  `json:"ability_score"`
	}
)
