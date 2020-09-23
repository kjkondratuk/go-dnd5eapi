// +build integration

package races

import (
	"net/http"
	"os"
	"testing"

	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/stretchr/testify/assert"
)

var (
	APIBaseURL = os.Getenv("API_ROOT")
	Client     = NewRaceClient(api.NewBasicsProvider(&http.Client{}, APIBaseURL))
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
