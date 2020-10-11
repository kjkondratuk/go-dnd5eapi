//go:generate go run ../gen/api_gen.go ../gen magic_schools MagicSchool MagicSchoolDetail "\"abjuration\""

package magic_schools

import (
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	MagicSchoolDetail struct {
		Index       string  `json:"index"`
		Name        string  `json:"name"`
		Url         string  `json:"url"`
		Class       api.Ref `json:"class"`
		Level       int     `json:"level"`
		Description string  `json:"desc"`
	}
)
