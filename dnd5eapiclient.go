package go_dnd5eapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

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

// region module-private functions
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
