package go_dnd5eapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kjkondratuk/go-dnd5eapi/response"
)

type apiClient struct {
	client  *http.Client
	baseURL string
}

type ApiClient interface {
	GetAllEndpoints() (*response.EndpointResponse, error)
	GetAllAbilityScores() (*response.AllAbilityScoresResponse, error)
	GetAbilityScoreByName(name string) (*response.AbilityScoreDetail, error)
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
func (ac *apiClient) GetAllEndpoints() (*response.EndpointResponse, error) {
	result, err := ac.apiGet("")
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
func (ac *apiClient) GetAllAbilityScores() (*response.AllAbilityScoresResponse, error) {
	result, err := ac.apiGet("/ability-scores")
	if err != nil {
		return nil, err
	}

	d := response.AllAbilityScoresResponse{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (ac *apiClient) GetAbilityScoreByName(name string) (*response.AbilityScoreDetail, error) {
	result, err := ac.apiGet("/ability-scores/" + strings.ToLower(name))
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
