// This file is generated by go-dnd5eapi/gen/api_gen.go.  DO NOT EDIT THIS FILE.  Generation parameters are:
// Package Name: languages
// Endpoint: languages
// API Name: Language
// LC API Name: language
// API Class: LanguageDetail

package languages

import (
	json "github.com/json-iterator/go"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

type (
	languageClient struct {
		api.BasicsProvider
		uri string
	}
)

type LanguageClient interface {
	GetList() (*api.ListResponse, error)
	QueryList(query map[string]string) (*api.ListResponse, error)
	GetByIndex(index string) (*LanguageDetail, error)
}

func NewClient(basicsProvider api.BasicsProvider) LanguageClient {
	return &languageClient{
		BasicsProvider: basicsProvider,
		uri: "/languages",
	}
}

func (ac *languageClient) GetList() (*api.ListResponse, error) {
	return ac.GetListForUrl(ac.uri)
}

func (ac *languageClient) QueryList(query map[string]string) (*api.ListResponse, error) {
	return ac.QueryListForUrl(ac.uri, query)
}

func (ac *languageClient) GetByIndex(index string) (*LanguageDetail, error) {
	result, err := ac.ApiGet(ac.uri + "/" + index)
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
