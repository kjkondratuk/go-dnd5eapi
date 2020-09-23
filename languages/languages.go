package languages

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for languages
	LanguagesEndpoint      = "/languages"
	LanguagesChildEndpoint = LanguagesEndpoint + "/"
)

type (
	LanguageChoice struct {
		Choose int          `json:"choose"`
		Type   string       `json:"type"`
		From   []api.APIRef `json:"from"`
	}

	LanguageDetail struct {
		Index           string   `json:"index"`
		Name            string   `json:"name"`
		Url             string   `json:"url"`
		Type            string   `json:"type"`
		TypicalSpeakers []string `json:"typical_speakers"`
		Script          string   `json:"script"`
	}

	languageClient struct {
		basicsProvider api.BasicsProvider
	}

	LanguageClient interface {
		GetLanguageList() (*api.ListResponse, error)
		GetLanguageByIndex(index string) (*LanguageDetail, error)
	}
)

func NewLanguagesClient(basicsProvider api.BasicsProvider) LanguageClient {
	return &languageClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *languageClient) GetLanguageList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(LanguagesEndpoint)
}

func (ac *languageClient) GetLanguageByIndex(index string) (*LanguageDetail, error) {
	result, err := ac.basicsProvider.ApiGet(LanguagesChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := LanguageDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
