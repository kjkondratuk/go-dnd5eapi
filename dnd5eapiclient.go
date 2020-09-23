package go_dnd5eapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type (
	apiClient struct {
		client  *http.Client
		baseURL string
	}

	EndpointResponse map[string]string

	APIRef struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	}

	ListResponse struct {
		Count   int `json:"count"`
		Results []APIRef
	}

	ApiClient interface {
		// Endpoints
		GetEndpointList() (*EndpointResponse, error)

		// Ability Score
		GetAbilityScoreList() (*ListResponse, error)
		GetAbilityScoreByIndex(index string) (*AbilityScoreDetail, error)

		// Skills
		GetSkillList() (*ListResponse, error)
		GetSkillByIndex(index string) (*SkillDetail, error)

		// Skill <-> Ability Score
		GetAbilityScoreForSkillByIndex(skillIndex string) (*AbilityScoreDetail, error)
		GetSkillsForAbilityByIndex(index string) ([]SkillDetail, error)

		// Class
		GetClassList() (*ListResponse, error)
		GetClassByIndex(index string) (*ClassDetail, error)

		// Class <-> Proficiencies
		GetProficienciesForClassByIndex(classIndex string) ([]ProficiencyDetail, error)
		GetProficiencyChoicesForClassByIndex(classIndex string) (map[string][]ProficiencyDetail, error)

		// Proficiencies
		GetProficiencyList() (*ListResponse, error)
		GetProficiencyByIndex(index string) (*ProficiencyDetail, error)

		// Subclasses
		GetSubclassList() (*ListResponse, error)

		// Races
		GetRaceList() (*ListResponse, error)
		GetRaceByIndex(index string) (*RaceDetail, error)

		// Languages
		GetLanguageList() (*ListResponse, error)
		GetLanguageByIndex(index string) (*LanguageDetail, error)
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

func (ac *apiClient) getListForUrl(url string) (*ListResponse, error) {
	result, err := ac.apiGet(url)
	if err != nil {
		return nil, err
	}

	d := ListResponse{}
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
