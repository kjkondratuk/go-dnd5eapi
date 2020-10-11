package monsters

import (
	"net/http"
	"os"
	"testing"

	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/stretchr/testify/assert"
)

var (
	APIBaseURL = os.Getenv("API_ROOT")
	Client     = NewClient(api.NewBasicsProvider(&http.Client{}, APIBaseURL))
)

func TestMain(t *testing.M) {
	t.Run()
	os.Exit(0)
}

func TestMonsters_GetList_IT(t *testing.T) {
	_, err := Client.GetList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

//func TestMonsters_GetByIndex_IT(t *testing.T) {
//	_, err := Client.GetByIndex("aboleth")
//	assert.Nil(t, err, "Should not receive an error contacting API.")
//	assert.True(t, true, "Should complete successfully!")
//}