package go_dnd5eapi

import (
	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/kjkondratuk/go-dnd5eapi/api/ability_scores"
	"github.com/kjkondratuk/go-dnd5eapi/classes"
	classes_proficiency_rel "github.com/kjkondratuk/go-dnd5eapi/classes_proficiencies_rel"
	"github.com/kjkondratuk/go-dnd5eapi/conditions"
	"github.com/kjkondratuk/go-dnd5eapi/damage_types"
	"github.com/kjkondratuk/go-dnd5eapi/endpoints"
	"github.com/kjkondratuk/go-dnd5eapi/equipment"
	"github.com/kjkondratuk/go-dnd5eapi/features"
	"github.com/kjkondratuk/go-dnd5eapi/languages"
	"github.com/kjkondratuk/go-dnd5eapi/magic_schools"
	"github.com/kjkondratuk/go-dnd5eapi/monsters"
	"github.com/kjkondratuk/go-dnd5eapi/proficiencies"
	"github.com/kjkondratuk/go-dnd5eapi/races"
	"github.com/kjkondratuk/go-dnd5eapi/skills"
	"github.com/kjkondratuk/go-dnd5eapi/skills_ability_rel"
	"github.com/kjkondratuk/go-dnd5eapi/spellcasting"
	"github.com/kjkondratuk/go-dnd5eapi/spells"
	"github.com/kjkondratuk/go-dnd5eapi/starting_equipment"
	"github.com/kjkondratuk/go-dnd5eapi/subclasses"
	"github.com/kjkondratuk/go-dnd5eapi/subraces"
	"github.com/kjkondratuk/go-dnd5eapi/traits"
	"github.com/kjkondratuk/go-dnd5eapi/weapon_properties"
)

type (
	Client struct {
		basicsProvider     api.BasicsProvider
		AbilityScores      ability_scores.AbilityScoreClient
		Classes            classes.ClassClient
		ClassProficiencies classes_proficiency_rel.ClassProficiencyRelClient
		Conditions         conditions.ConditionClient
		Damage             damage_types.DamageTypeClient
		Endpoints          endpoints.EndpointsClient
		Equipment          equipment.EquipmentClient
		Features           features.FeatureClient
		Languages          languages.LanguageClient
		MagicSchools       magic_schools.MagicSchoolClient
		Monsters           monsters.MonsterClient
		Proficiencies      proficiencies.ProficiencyClient
		Races              races.RaceClient
		Skills             skills.SkillClient
		SkillsAbility      skills_ability_rel.SkillAbilityRelClient
		Spellcasting       spellcasting.SpellcastingClient
		Spells             spells.SpellClient
		StartingEquipment  starting_equipment.StartingEquipmentClient
		Subclasses         subclasses.SubclassClient
		Subraces           subraces.SubraceClient
		Traits             traits.TraitClient
		WeaponProperties   weapon_properties.WeaponPropertiesClient
	}
)

func NewClient(basicsProvider api.BasicsProvider) *Client {
	skillsClient := skills.NewClient(basicsProvider)
	abilityScoresClient := ability_scores.NewClient(basicsProvider)
	classesClient := classes.NewClient(basicsProvider)
	proficienciesClient := proficiencies.NewClient(basicsProvider)
	client := &Client{
		basicsProvider:     basicsProvider,
		AbilityScores:      abilityScoresClient,
		Classes:            classesClient,
		ClassProficiencies: classes_proficiency_rel.NewClient(classesClient, proficienciesClient),
		Conditions:         conditions.NewClient(basicsProvider),
		Damage:             damage_types.NewClient(basicsProvider),
		Endpoints:          endpoints.NewClient(basicsProvider),
		Equipment:          equipment.NewClient(basicsProvider),
		Features:           features.NewClient(basicsProvider),
		Languages:          languages.NewClient(basicsProvider),
		MagicSchools:       magic_schools.NewClient(basicsProvider),
		Monsters:           monsters.NewClient(basicsProvider),
		Proficiencies:      proficienciesClient,
		Races:              races.NewClient(basicsProvider),
		Skills:             skillsClient,
		SkillsAbility:      skills_ability_rel.NewClient(skillsClient, abilityScoresClient),
		Spellcasting:       spellcasting.NewClient(basicsProvider),
		Spells:             spells.NewClient(basicsProvider),
		StartingEquipment:  starting_equipment.NewClient(basicsProvider),
		Subclasses:         subclasses.NewClient(basicsProvider),
		Subraces:           subraces.NewClient(basicsProvider),
		Traits:             traits.NewClient(basicsProvider),
		WeaponProperties:   weapon_properties.NewClient(basicsProvider),
	}
	return client
}
