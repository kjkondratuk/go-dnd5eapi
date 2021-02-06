// +build integration

package skills_ability_rel

import (
	"net/http"
	"os"
	"testing"

	"github.com/kjkondratuk/go-dnd5eapi/ability_scores"
	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/kjkondratuk/go-dnd5eapi/skills"
	"github.com/stretchr/testify/assert"
)

var (
	APIBaseURL = os.Getenv("API_ROOT")
	base       = api.NewBasicsProvider(&http.Client{}, APIBaseURL)
	Client     = NewClient(skills.NewClient(base), ability_scores.NewClient(base))
)

func TestSkillAbilityRel_GetSkillsForAbilityByIndex(t *testing.T) {
	_, err := Client.GetSkillsForAbilityByIndex("cha")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

func TestSkillAbilityRel_GetAbilityScoreForSkillByIndex(t *testing.T) {
	_, err := Client.GetAbilityScoreForSkillByIndex("cha")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}
