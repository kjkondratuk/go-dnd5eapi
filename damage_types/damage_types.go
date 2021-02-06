//go:generate go run ../gen/api_gen.go DamageType "\"lightning\""

package damage_types

import "github.com/kjkondratuk/go-dnd5eapi/api"

type DamageTypeDetail struct {
	api.Description
}
