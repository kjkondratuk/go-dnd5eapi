// +build integration

package classes

import (
	"net/http"
	"os"
	"testing"

	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/stretchr/testify/assert"
)

var (
	APIBaseURL = os.Getenv("API_ROOT")
	Client     = NewClassesClient(api.NewBasicsProvider(&http.Client{}, APIBaseURL))
)

func TestClasses_GetClassByIndex_IT(t *testing.T) {
	_, err := Client.GetClassByIndex("bard")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}
