//go:generate go run ../gen/api_gen.go weapon_properties WeaponProperties "\"ammunition\""

package weapon_properties

import "github.com/kjkondratuk/go-dnd5eapi/api"

type WeaponPropertiesDetail struct {
	api.Description
}
