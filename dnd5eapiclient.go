package go_dnd5eapi

import (
	"github.com/kjkondratuk/go-dnd5eapi/ability_scores"
	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/kjkondratuk/go-dnd5eapi/classes"
	classes_proficiency_rel "github.com/kjkondratuk/go-dnd5eapi/classes_proficiencies_rel"
	"github.com/kjkondratuk/go-dnd5eapi/conditions"
	"github.com/kjkondratuk/go-dnd5eapi/damage"
	"github.com/kjkondratuk/go-dnd5eapi/endpoints"
	"github.com/kjkondratuk/go-dnd5eapi/equipment"
	"github.com/kjkondratuk/go-dnd5eapi/features"
	"github.com/kjkondratuk/go-dnd5eapi/languages"
	"github.com/kjkondratuk/go-dnd5eapi/magic_schools"
	"github.com/kjkondratuk/go-dnd5eapi/monsters"
	"github.com/kjkondratuk/go-dnd5eapi/proficiencies"
	"github.com/kjkondratuk/go-dnd5eapi/races"
	"github.com/kjkondratuk/go-dnd5eapi/skills"
	"github.com/kjkondratuk/go-dnd5eapi/subclasses"
	"github.com/kjkondratuk/go-dnd5eapi/traits"
)

type (
	Client struct {
		basicsProvider     api.BasicsProvider
		AbilityScores      ability_scores.AbilityScoreClient
		Classes            classes.ClassClient
		ClassProficiencies classes_proficiency_rel.ClassProficiencyRelClient
		Conditions         conditions.ConditionClient
		Damage             damage.DamageTypeClient
		Endpoints          endpoints.EndpointsClient
		Equipment          equipment.EquipmentClient
		Features           features.FeatureClient
		Languages          languages.LanguageClient
		MagicSchools       magic_schools.MagicSchoolClient
		Monsters           monsters.MonsterClient
		Proficiencies      proficiencies.ProficiencyClient
		Races              races.RaceClient
		Skills             skills.SkillClient
		Subclasses         subclasses.SubclassClient
		Traits             traits.TraitClient
	}
)

func NewClient(basicsProvider api.BasicsProvider) *Client {
	client := &Client{
		basicsProvider:     basicsProvider,
		AbilityScores:      ability_scores.NewClient(basicsProvider),
		Classes:            classes.NewClient(basicsProvider),
		ClassProficiencies: nil,
		Conditions:         conditions.NewClient(basicsProvider),
		Damage:             damage.NewClient(basicsProvider),
		Endpoints:          endpoints.NewClient(basicsProvider),
		Equipment:          equipment.NewClient(basicsProvider),
		Features:           features.NewClient(basicsProvider),
		Languages:          languages.NewClient(basicsProvider),
		MagicSchools:       magic_schools.NewClient(basicsProvider),
		Monsters:           monsters.NewClient(basicsProvider),
		Proficiencies:      proficiencies.NewClient(basicsProvider),
		Races:              races.NewClient(basicsProvider),
		Skills:             skills.NewClient(basicsProvider),
		Subclasses:         subclasses.NewClient(basicsProvider),
		Traits:             traits.NewClient(basicsProvider),
	}
	client.ClassProficiencies = classes_proficiency_rel.NewClient(client.Classes, client.Proficiencies)
	return client
}
