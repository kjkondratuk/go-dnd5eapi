//go:generate go run ../gen/api_gen.go conditions Condition "\"blinded\""

package conditions

import "github.com/kjkondratuk/go-dnd5eapi/api"

type ConditionDetail struct {
	api.Description
}
