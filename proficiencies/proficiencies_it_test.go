// +build integration

package proficiencies

import (
	"net/http"
	"os"
	"testing"

	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/stretchr/testify/assert"
)

var (
	APIBaseURL = os.Getenv("API_ROOT")
	Client     = NewProficiencyClient(api.NewBasicsProvider(&http.Client{}, APIBaseURL))
)

func TestProficiencies_GetProficiencyList_IT(t *testing.T) {
	_, err := Client.GetProficiencyList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

func TestProficiencies_GetProficiencyByIndex_IT(t *testing.T) {
	_, err := Client.GetProficiencyByIndex("battleaxes")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}
