// +build integration

package go_dnd5eapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClassProficiencyRel_GetProficienciesForClassByIndex_IT(t *testing.T) {
	_, err := Client.GetProficienciesForClassByIndex("bard")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

func TestClassProficiencyRel_GetProficiencyChoicesForClassByIndex_IT(t *testing.T) {
	_, err := Client.GetProficiencyChoicesForClassByIndex("bard")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

