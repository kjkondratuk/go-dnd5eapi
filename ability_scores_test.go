// +build integration

package go_dnd5eapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbilityScores_GetAbilityScoreList(t *testing.T) {
	_, err := Client.GetAbilityScoreList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

func TestAbilityScores_GetAbilityScoreByName(t *testing.T) {
	_, err := Client.GetAbilityScoreByIndex("CHA")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}