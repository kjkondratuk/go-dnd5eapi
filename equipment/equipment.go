//go:generate go run ../gen/api_gen.go ../gen equipment Equipment EquipmentDetail "\"abacus\""

package equipment

import (
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	EquipmentDetail struct {
		Index             string       `json:"index"`
		Name              string       `json:"name"`
		Url               string       `json:"url"`
		EquipmentCategory api.Ref      `json:"equipment_category"`
		WeaponCategory    string       `json:"weapon_category"`
		WeaponRange       string       `json:"weapon_range"`
		CategoryRange     string       `json:"category_range"`
		Cost              api.Cost     `json:"cost"`
		Damage            DamageDetail `json:"damage"`
		Range             api.Range    `json:"range"`
		Weight            int          `json:"weight"`
		Properties        []api.Ref    `json:"properties"`
	}

	DamageDetail struct {
		DamageDice string
		DamageType api.Ref
	}
)
