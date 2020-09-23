// +build integration

package languages

import (
	"net/http"
	"os"
	"testing"

	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/stretchr/testify/assert"
)

var (
	APIBaseURL = os.Getenv("API_ROOT")
	Client     = NewLanguagesClient(api.NewBasicsProvider(&http.Client{}, APIBaseURL))
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
