//go:generate go run ../gen/api_gen.go Monster "\"aboleth\""

package monsters

import (
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	MonsterDetail struct {
		Index                 string             `json:"index"`
		Name                  string             `json:"name"`
		Url                   string             `json:"url"`
		Size                  string             `json:"size"`
		Type                  string             `json:"type"`
		Subtype               string             `json:"subtype"`
		Alignment             string             `json:"alignment"`
		ArmorClass            int                `json:"armor_class"`
		HitPoints             int                `json:"hit_points"`
		HitDice               string             `json:"hit_dice"`
		Speed                 Speed              `json:"speed"`
		Strength              int                `json:"strength"`
		Dexterity             int                `json:"dexterity"`
		Constitution          int                `json:"constitution"`
		Intelligence          int                `json:"intelligence"`
		Wisdom                int                `json:"wisdom"`
		Charisma              int                `json:"charisma"`
		Proficiencies         []ProficiencyValue `json:"proficiencies"`
		DamageVulnerabilities []string           `json:"damage_vulnerabilities"`
		DamageResistances     []string           `json:"damage_resistances"`
		DamageImmunities      []string           `json:"damage_immunities"`
		ConditionImmunities   []string           `json:"condition_immunities"`
		Senses                Sense              `json:"senses"`
		Languages             string             `json:"languages"`
		ChallengeRating       int                `json:"challenge_rating"`
		SpecialAbilities      []SpecialAbility   `json:"special_abilities"`
		Actions               []Action           `json:"actions"`
		LegendaryActions      []Action           `json:"legendary_actions"`
	}

	Action struct {
		Name        string     `json:"name"`
		Description string     `json:"desc"`
		Options     api.Choice `json:"options"`
		AttackBonus *int       `json:"attack_bonus,omitempty"`
		DC          *DC        `json:"dc,omitempty"`
		Usage       *Usage     `json:"usage,omitempty"`
		Damage      []Damage   `json:"damage_types"`
	}

	Damage struct {
		Type api.Ref `json:"damage_type"`
		Dice string  `json:"damage_dice"`
	}

	SpecialAbility struct {
		Name        string `json:"name"`
		Description string `json:"desc"`
		Usage       *Usage `json:"usage,omitempty"`
		DC          *DC    `json:"dc,omitempty"`
	}

	DC struct {
		Type        api.Ref `json:"dc_type"`
		Value       int     `json:"dc_value"`
		SuccessType string  `json:"success_type"`
	}

	Usage struct {
		Type  string `json:"type"`
		Times int    `json:"times"`
	}

	Sense struct {
		Blindsight        string `json:"blindsight"`
		Darkvision        string `json:"darkvision"`
		PassivePerception int    `json:"passive_perception"`
	}

	ProficiencyValue struct {
		Value       int     `json:"value"`
		Proficiency api.Ref `json:"proficiency"`
	}

	Speed struct {
		Walk  string `json:"walk"`
		Climb string `json:"climb"`
		Fly   string `json:"fly"`
	}
)
