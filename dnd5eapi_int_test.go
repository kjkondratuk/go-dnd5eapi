// +build integration

package go_dnd5eapi

const (
	APIBaseURL = "https://www.dnd5eapi.co/api"
)

var (
	Client = NewApiClient(APIBaseURL)
)