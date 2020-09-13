package go_dnd5eapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/kjkondratuk/go-dnd5eapi/constants"
	"github.com/kjkondratuk/go-dnd5eapi/response"
)

type (
	apiClient struct {
		client  *http.Client
		baseURL string
	}

	ApiClient interface {
		// Endpoints
		GetEndpointList() (*response.EndpointResponse, error)

		// Ability Score
		GetAbilityScoreList() (*response.ListResponse, error)
		GetAbilityScoreByIndex(index string) (*response.AbilityScoreDetail, error)

		// Skills
		GetSkillList() (*response.ListResponse, error)
		GetSkillByIndex(index string) (*response.SkillDetail, error)

		// Skill <-> Ability Score
		GetAbilityScoreForSkillByIndex(skillIndex string) (*response.AbilityScoreDetail, error)
		GetSkillsForAbilityByIndex(index string) ([]response.SkillDetail, error)

		// Class
		GetClassList() (*response.ListResponse, error)
		GetClassByIndex(index string) (*response.ClassDetail, error)

		// Class <-> Proficiencies
		GetProficienciesForClassByIndex(classIndex string) ([]response.ProficiencyDetail, error)
	}
)

// region ApiClient Implementation
func NewApiClient(baseURL string, opts ...ClientOption) ApiClient {
	client := &apiClient{
		client:  &http.Client{},
		baseURL: baseURL,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// region private functions
func (ac *apiClient) apiGet(uri string) ([]byte, error) {
	resp, err := ac.client.Get(ac.baseURL + uri)
	if err != nil {
		return nil, err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	} else {
		return nil, errors.New("response was empty")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ac *apiClient) getListForUrl(url string) (*response.ListResponse, error) {
	result, err := ac.apiGet(url)
	if err != nil {
		return nil, err
	}

	d := response.ListResponse{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// endregion

// endregion

// region Client Options
type ClientOption func(*apiClient)

func WithHttpClient(client *http.Client) ClientOption {
	return func(ac *apiClient) {
		ac.client = client
	}
}

// endregion

// region Endpoints
func (ac *apiClient) GetEndpointList() (*response.EndpointResponse, error) {
	result, err := ac.apiGet(constants.RootEndpoint)
	if err != nil {
		return nil, err
	}

	d := response.EndpointResponse{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// endregion

// region Ability Scores
func (ac *apiClient) GetAbilityScoreList() (*response.ListResponse, error) {
	return ac.getListForUrl(constants.AbilityScoreEndpoint)
}

func (ac *apiClient) GetAbilityScoreByIndex(index string) (*response.AbilityScoreDetail, error) {
	result, err := ac.apiGet(constants.AbilityScoreChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := response.AbilityScoreDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// endregion

// region Skills

func (ac *apiClient) GetSkillList() (*response.ListResponse, error) {
	return ac.getListForUrl(constants.SkillsEndpoint)
}

func (ac *apiClient) GetSkillByIndex(index string) (*response.SkillDetail, error) {
	result, err := ac.apiGet(constants.SkillsChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := response.SkillDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// endregion

// region Skill <-> Ability Score
func (ac *apiClient) GetSkillsForAbilityByIndex(index string) ([]response.SkillDetail, error) {
	ability, err := ac.GetAbilityScoreByIndex(index)
	if err != nil {
		return nil, err
	}

	var result []response.SkillDetail
	for _, v := range ability.Skills {
		skill, err := ac.GetSkillByIndex(v.Index)
		if err != nil {
			return nil, err
		}
		if skill != nil {
			result = append(result, *skill)
		} else {
			return nil, errors.New("No skill called " + v.Index + " could be found")
		}
	}
	return result, nil
}

func (ac *apiClient) GetAbilityScoreForSkillByIndex(skillIndex string) (*response.AbilityScoreDetail, error) {
	skill, err := ac.GetSkillByIndex(skillIndex)
	if err != nil {
		return nil, err
	}
	abilityScore, err := ac.GetAbilityScoreByIndex(skill.AbilityScore.Index)
	if err != nil {
		return nil, err
	}
	return abilityScore, nil
}

// endregion

// region Classes
func (ac *apiClient) GetClassList() (*response.ListResponse, error) {
	return ac.getListForUrl(constants.ClassEndpoint)
}

func (ac *apiClient) GetClassByIndex(index string) (*response.ClassDetail, error) {
	result, err := ac.apiGet(constants.ClassChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := response.ClassDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// endregion

// region Classes <-> Proficiency
func (ac *apiClient) GetProficienciesForClassByIndex(classIndex string) ([]response.ProficiencyDetail, error) {
	class, err := ac.GetClassByIndex(classIndex)
	if err != nil {
		return nil, err
	}

	var result []response.ProficiencyDetail
	for _, v := range class.Proficiencies {
		prof, err := ac.GetProficiencyByIndex(v.Index)
		if err != nil {
			return nil, err
		}
		if prof != nil {
			result = append(result, *prof)
		} else {
			return nil, errors.New("No proficiency called " + v.Index + " could be found")
		}
	}
	return result, nil
}

func (ac *apiClient) GetProficiencyChoicesForClassByIndex(classIndex string) (map[string][]response.ProficiencyDetail, error) {
	class, err := ac.GetClassByIndex(classIndex)
	if err != nil {
		return nil, err
	}

	var result = make(map[string][]response.ProficiencyDetail)
	for i, choice := range class.ProficiencyChoices {
		var profs []response.ProficiencyDetail
		for _, prof := range class.ProficiencyChoices[i].From {
			p, err := ac.GetProficiencyByIndex(prof.Index)
			if err != nil {
				return nil, err
			}
			if p != nil {
				profs = append(profs, *p)
			} else {
				return nil, errors.New("No proficiency called " + prof.Index + " could be found")
			}
		}
		result[choice.Type] = profs
	}
	return result, nil
}

// endregion

//region Proficiencies
func (ac *apiClient) GetProficiencyByIndex(index string) (*response.ProficiencyDetail, error) {
	result, err := ac.apiGet(constants.ProficienciesChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := response.ProficiencyDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// endregion

// region Subclasses
func (ac *apiClient) GetSubclassList() (*response.ListResponse, error) {
	return ac.getListForUrl(constants.SubclassesEndpoint)
}

// endregion
