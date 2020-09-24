package go_dnd5eapi

import (
	"github.com/kjkondratuk/go-dnd5eapi/ability_scores"
	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/kjkondratuk/go-dnd5eapi/classes"
	"github.com/kjkondratuk/go-dnd5eapi/endpoints"
	"github.com/kjkondratuk/go-dnd5eapi/languages"
	"github.com/kjkondratuk/go-dnd5eapi/proficiencies"
	"github.com/kjkondratuk/go-dnd5eapi/races"
	"github.com/kjkondratuk/go-dnd5eapi/skills"
	"github.com/kjkondratuk/go-dnd5eapi/subclasses"
)

type (
	Client struct {
		basicsProvider api.BasicsProvider
		AbilityScores  ability_scores.AbilityScoreClient
		Classes        classes.ClassesClient
		Endpoints      endpoints.EndpointsClient
		Languages      languages.LanguageClient
		Proficiencies  proficiencies.ProficiencyClient
		Races          races.RaceClient
		Skills         skills.SkillClient
		Subclasses     subclasses.SubclassClient
	}
)

func NewClient(basicsProvider api.BasicsProvider) *Client {
	return &Client{
		basicsProvider: basicsProvider,
		AbilityScores:  ability_scores.NewClient(basicsProvider),
		Classes:        classes.NewClient(basicsProvider),
		Endpoints:      endpoints.NewClient(basicsProvider),
		Languages:      languages.NewClient(basicsProvider),
		Proficiencies:  proficiencies.NewClient(basicsProvider),
		Races:          races.NewClient(basicsProvider),
		Skills:         skills.NewClient(basicsProvider),
		Subclasses:     subclasses.NewClient(basicsProvider),
	}
}
