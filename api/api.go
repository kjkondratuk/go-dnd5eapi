package api

import (
	"errors"
	"io/ioutil"
	"net/http"

	json "github.com/json-iterator/go"
)

type (
	ListResponse struct {
		Count   int `json:"count"`
		Results []Ref
	}

	Ref struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	}

	Description struct {
		Index       *string  `json:"index"`
		Name        string   `json:"name"`
		Url         *string  `json:"url"`
		Description []string `json:"desc"`
	}

	Range struct {
		Normal int
		Long   int
	}

	Cost struct {
		Quantity int
		Unit     string
	}

	Choice struct {
		Choose int           `json:"choose"`
		Type   string        `json:"type"`
		From   []interface{} `json:"from"`
	}

	basicsProvider struct {
		client  *http.Client
		baseUrl string
	}

	BasicsProvider interface {
		ApiGet(uri string) ([]byte, error)
		GetListForUrl(url string) (*ListResponse, error)
	}
)

func NewBasicsProvider(client *http.Client, baseUrl string) BasicsProvider {
	return &basicsProvider{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (bp *basicsProvider) ApiGet(uri string) ([]byte, error) {
	resp, err := bp.client.Get(bp.baseUrl + uri)
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

func (bp *basicsProvider) GetListForUrl(uri string) (*ListResponse, error) {
	result, err := bp.ApiGet(uri)
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
