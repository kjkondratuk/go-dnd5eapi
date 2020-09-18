// +build integration

package go_dnd5eapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRaces_GetRaceList_IT(t *testing.T) {
	_, err := Client.GetRaceList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

func TestRaces_GetRaceByIndex_IT(t *testing.T) {
	_, err := Client.GetRaceByIndex("half-orc")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}
