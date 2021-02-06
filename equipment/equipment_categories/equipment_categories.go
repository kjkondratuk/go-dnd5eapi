//go:generate go run ../../gen/api_gen.go Category "\"shields\""

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
