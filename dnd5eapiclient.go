package go_dnd5eapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kjkondratuk/go-dnd5eapi/constants"
	"github.com/kjkondratuk/go-dnd5eapi/response"
)

type apiClient struct {
	client  *http.Client
	baseURL string
}

type ApiClient interface {
	GetEndpointList() (*response.EndpointResponse, error)
	GetAbilityScoreList() (*response.ListResponse, error)
	GetAbilityScoreByName(name string) (*response.AbilityScoreDetail, error)
	GetSkillList() (*response.ListResponse, error)
	GetSkillByName(name string) (*response.SkillDetail, error)
}

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

// endregion

// region Options
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
	result, err := ac.apiGet(constants.AbilityScoreEndpoint)
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

func (ac *apiClient) GetAbilityScoreByName(name string) (*response.AbilityScoreDetail, error) {
	result, err := ac.apiGet(constants.AbilityScoreChildEndpoint + strings.ToLower(name))
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
	result, err := ac.apiGet(constants.SkillsEndpoint)
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

func (ac * apiClient) GetSkillByName(name string) (*response.SkillDetail, error) {
	result, err := ac.apiGet(constants.SkillsChildEndpoint + strings.ReplaceAll(strings.ToLower(name), " ", "-"))
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

func (ac *apiClient) GetAbilityScoreForSkill(skillName string) (*response.AbilityScoreDetail, error) {
	skill, err := ac.GetSkillByName(skillName)
	if err != nil {
		return nil, err
	}
	abilityScore, err := ac.GetAbilityScoreByName(skill.AbilityScore.Name)
	if err != nil {
		return nil, err
	}
	return abilityScore, nil
}

// endregion