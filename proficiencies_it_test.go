// +build integration

package go_dnd5eapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProficiencies_GetProficiencyByIndex(t *testing.T) {
	_, err := Client.GetProficiencyByIndex("battleaxes")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}