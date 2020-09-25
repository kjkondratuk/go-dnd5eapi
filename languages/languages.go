package languages

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for languages
	Endpoint      = "/languages"
	ChildEndpoint = Endpoint + "/"
)

type (
	LanguageChoice struct {
		Choose int       `json:"choose"`
		Type   string    `json:"type"`
		From   []api.Ref `json:"from"`
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
		GetList() (*api.ListResponse, error)
		GetByIndex(index string) (*LanguageDetail, error)
	}
)

func NewClient(basicsProvider api.BasicsProvider) LanguageClient {
	return &languageClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *languageClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *languageClient) GetByIndex(index string) (*LanguageDetail, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEndpoint + index)
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
