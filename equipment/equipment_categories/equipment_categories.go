//go:generate go run ../../gen/api_gen.go ../../gen equipment_categories Category CategoryDetail "\"shields\""

package equipment_categories

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
