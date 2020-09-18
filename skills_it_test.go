// +build integration

package go_dnd5eapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkills_GetSkillList_IT(t *testing.T) {
	_, err := Client.GetSkillList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

func TestSkills_GetSkillByIndex_IT(t *testing.T) {
	_, err := Client.GetSkillByIndex("acrobatics")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}
