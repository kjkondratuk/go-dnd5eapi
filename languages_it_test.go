// +build integration

package go_dnd5eapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRaces_GetLanguageList_IT(t *testing.T) {
	_, err := Client.GetLanguageList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

func TestRaces_GetLanguageByIndex_IT(t *testing.T) {
	_, err := Client.GetLanguageByIndex("abyssal")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}
