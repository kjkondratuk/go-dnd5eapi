//go:generate go run ../gen/api_gen.go starting_equipment StartingEquipment "\"barbarian\""

package starting_equipment

import "github.com/kjkondratuk/go-dnd5eapi/api"

type (
	StartingEquipmentDetail struct {
		Index                    string          `json:"index"`
		Name                     string          `json:"name"`
		Url                      string          `json:"url"`
		Class                    api.Ref         `json:"class"`
		StartingEquipment        []EquipmentItem `json:"starting_equipment"`
		StartingEquipmentOptions []api.Choice    `json:"starting_equipment_options"`
	}

	EquipmentItem struct {
		Equipment api.Ref `json:"equipment"`
		Quantity  int     `json:"quantity"`
	}
)
