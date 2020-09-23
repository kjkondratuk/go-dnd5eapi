// +build integration

package subclasses

import (
	"net/http"
	"os"
	"testing"

	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/stretchr/testify/assert"
)

var (
	APIBaseURL = os.Getenv("API_ROOT")
	Client     = NewSubclassClient(api.NewBasicsProvider(&http.Client{}, APIBaseURL))
)

func TestSubclasses_GetSubclassList_IT(t *testing.T) {
	_, err := Client.GetSubclassList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

func TestSubclasses_GetSubclassByIndex_IT(t *testing.T) {
	_, err := Client.GetSubclassByIndex("berserker")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}
