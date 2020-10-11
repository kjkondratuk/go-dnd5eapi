//go:generate go run ../../gen/api_gen.go ../../gen categories Category CategoryDetail "\"shields\""

package categories

import (
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	CategoryDetail struct {
		Index     string    `json:"index"`
		Name      string    `json:"name"`
		Url       string    `json:"url"`
		Equipment []api.Ref `json:"equipment"`
	}
)
