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
		asc            ability_scores.AbilityScoreClient
		cc             classes.ClassesClient
		ec             endpoints.EndpointsClient
		lc             languages.LanguageClient
		pc             proficiencies.ProficiencyClient
		rc             races.RaceClient
		sc             skills.SkillClient
		scc            subclasses.SubclassClient
	}
)

func NewClient(basicsProvider api.BasicsProvider) *Client {
	return &Client{
		basicsProvider: basicsProvider,
		asc:            ability_scores.NewClient(basicsProvider),
		cc:             classes.NewClient(basicsProvider),
		ec:             endpoints.NewClient(basicsProvider),
		lc:             languages.NewClient(basicsProvider),
		pc:             proficiencies.NewClient(basicsProvider),
		rc:             races.NewClient(basicsProvider),
		sc:             skills.NewClient(basicsProvider),
		scc:            subclasses.NewClient(basicsProvider),
	}
}
