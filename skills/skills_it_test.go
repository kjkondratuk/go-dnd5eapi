// +build integration

package skills

import (
	"net/http"
	"os"
	"testing"

	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/stretchr/testify/assert"
)

var (
	APIBaseURL = os.Getenv("API_ROOT")
	Client     = NewSkillClient(api.NewBasicsProvider(&http.Client{}, APIBaseURL))
)

func TestSkills_GetSkillList_IT(t *testing.T) {
	_, err := Client.GetSkillList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

func TestSkills_GetSkillByIndex_IT(t *testing.T) {
	_, err := Client.GetSkillByIndex("acrobatics")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}
