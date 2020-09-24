// +build integration

package classes_proficiency_rel

import (
	"net/http"
	"os"
	"testing"

	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/kjkondratuk/go-dnd5eapi/classes"
	"github.com/kjkondratuk/go-dnd5eapi/proficiencies"
	"github.com/stretchr/testify/assert"
)

var (
	APIBaseURL = os.Getenv("API_ROOT")
	base       = api.NewBasicsProvider(&http.Client{}, APIBaseURL)
	Client     = NewClient(classes.NewClient(base), proficiencies.NewClient(base))
)

func TestClassProficiencyRel_GetProficienciesForClassByIndex_IT(t *testing.T) {
	_, err := Client.GetProficienciesForClassByIndex("bard")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

// func TestClassProficiencyRel_GetProficiencyChoicesForClassByIndex_IT(t *testing.T) {
// 	_, err := Client.GetProficiencyChoicesForClassByIndex("bard")
// 	assert.Nil(t, err, "Should not receive an error contacting API.")
// 	assert.True(t, true, "Should complete successfully!")
// }
