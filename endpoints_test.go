package go_dnd5eapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpoints_GetEndpointList(t *testing.T) {
	_, err := Client.GetEndpointList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}