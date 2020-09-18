// +build integration

package go_dnd5eapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubclasses_GetSubclassList_IT(t *testing.T) {
	_, err := Client.GetSubclassList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}
