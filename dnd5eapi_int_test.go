// +build integration

package go_dnd5eapi

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	APIBaseURL = "https://www.dnd5eapi.co/api"
)

var (
	client = NewApiClient(APIBaseURL)
)

func TestApiClient_GetEndpointList(t *testing.T) {
	response, err := client.GetEndpointList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	_, _ = fmt.Fprintln(os.Stdout, response)
	assert.True(t, true, "Should complete successfully!")
}

func TestApiClient_GetAbilityScoreList(t *testing.T) {
	response, err := client.GetAbilityScoreList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	_, _ = fmt.Fprintln(os.Stdout, response)
	assert.True(t, true, "Should complete successfully!")
}

func TestApiClient_GetAbilityScoreByName(t *testing.T) {
	response, err := client.GetAbilityScoreByName("CHA")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	_, _ = fmt.Fprintln(os.Stdout, response)
	assert.True(t, true, "Should complete successfully!")
}
