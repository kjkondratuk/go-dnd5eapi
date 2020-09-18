// +build integration

package go_dnd5eapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO: update integration tests with more targeted assertions than just "this completed without error"
func TestAbilityScores_GetAbilityScoreList_IT(t *testing.T) {
	_, err := Client.GetAbilityScoreList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

func TestAbilityScores_GetAbilityScoreByName_IT(t *testing.T) {
	_, err := Client.GetAbilityScoreByIndex("CHA")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}