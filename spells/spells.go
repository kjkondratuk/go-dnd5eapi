//go:generate go run ../gen/api_gen.go spells Spell "\"acid-arrow\""

package spells

import "github.com/kjkondratuk/go-dnd5eapi/api"

type (
	SpellDetail struct {
		Index         string        `json:"index"`
		Name          string        `json:"name"`
		Url           string        `json:"url"`
		Description   []string      `json:"desc"`
		HigherLevel   []string      `json:"higher_level"`
		Components    []string      `json:"components"`
		Range         string        `json:"range"`
		Material      string        `json:"material"`
		Ritual        bool          `json:"ritual"`
		Duration      string        `json:"duration"`
		Concentration bool          `json:"concentration"`
		CastingTime   string        `json:"casting_time"`
		Level         int           `json:"level"`
		AttackType    string        `json:"attack_type"`
		Damage        CastingDamage `json:"damage"`
		School        api.Ref       `json:"school"`
		Classes       []api.Ref     `json:"classes"`
		Subclasses    []api.Ref     `json:"subclasses"`
	}

	CastingDamage struct {
		DamageType   api.Ref        `json:"damage_type"`
		DamageAtSlot map[int]string `json:"damage_at_slot_level"`
	}
)
