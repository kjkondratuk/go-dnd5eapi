// +build integration

package go_dnd5eapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
