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
	apiClient struct {
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

	ApiClient interface{}
)

func NewApiClient(basicsProvider api.BasicsProvider) ApiClient {
	return &apiClient{
		basicsProvider: basicsProvider,
		asc:            ability_scores.NewAbilityScoreClient(basicsProvider),
		cc:             classes.NewClassesClient(basicsProvider),
		ec:             endpoints.NewEndpointsClient(basicsProvider),
		lc:             languages.NewLanguagesClient(basicsProvider),
		pc:             proficiencies.NewProficiencyClient(basicsProvider),
		rc:             races.NewRaceClient(basicsProvider),
		sc:             skills.NewSkillClient(basicsProvider),
		scc:            subclasses.NewSubclassClient(basicsProvider),
	}
}
