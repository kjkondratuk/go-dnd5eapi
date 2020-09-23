// +build integration

package endpoints

import (
	"net/http"
	"os"
	"testing"

	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/stretchr/testify/assert"
)

var (
	APIBaseURL = os.Getenv("API_ROOT")
	Client     = NewEndpointsClient(api.NewBasicsProvider(&http.Client{}, APIBaseURL))
)

func TestEndpoints_GetEndpointList(t *testing.T) {
	_, err := Client.GetEndpointList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}
