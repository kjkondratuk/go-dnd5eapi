// +build integration

package go_dnd5eapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProficiencies_GetProficiencyList_IT(t *testing.T) {
	_, err := Client.GetProficiencyList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

func TestProficiencies_GetProficiencyByIndex(t *testing.T) {
	_, err := Client.GetProficiencyByIndex("battleaxes")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}
