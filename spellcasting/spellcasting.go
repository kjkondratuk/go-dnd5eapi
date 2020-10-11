//go:generate go run ../gen/api_gen.go ../gen spellcasting Spellcasting SpellcastingDetail "\"bard\""

package spellcasting

import "github.com/kjkondratuk/go-dnd5eapi/api"

type (
	SpellcastingDetail struct {
		Index               string            `json:"index"`
		Name                string            `json:"name"`
		Url                 string            `json:"url"`
		Class               api.Ref           `json:"class"`
		Level               int               `json:"level"`
		SpellcastingAbility api.Ref           `json:"spellcasting_ability"`
		Info                []api.Description `json:"info"`
	}
)
