// +build integration

package ability_scores

import (
	"net/http"
	"os"
	"testing"

	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/stretchr/testify/assert"
)

var (
	APIBaseURL = os.Getenv("API_ROOT")
	Client     = NewAbilityScoreClient(api.NewBasicsProvider(&http.Client{}, APIBaseURL))
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
