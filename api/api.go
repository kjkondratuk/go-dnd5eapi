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
		HttpClient
		baseUrl string
	}

	BasicsProvider interface {
		ApiGet(uri string) ([]byte, error)
		GetListForUrl(endpoint string) (*ListResponse, error)
		QueryListForUrl(endpoint string, query map[string]string) (*ListResponse, error)
	}

	HttpClient interface {
		Get(url string) (*http.Response, error)
		Do(req *http.Request) (*http.Response, error)
	}
)

func NewBasicsProvider(client HttpClient, baseUrl string) BasicsProvider {
	return &basicsProvider{
		HttpClient: client,
		baseUrl:    baseUrl,
	}
}

func (bp *basicsProvider) ApiGet(uri string) ([]byte, error) {
	resp, err := bp.Get(bp.baseUrl + uri)
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

	r, err := unmarshalListResponse(result)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (bp *basicsProvider) QueryListForUrl(uri string, query map[string]string) (*ListResponse, error) {
	req, err := http.NewRequest(http.MethodGet, bp.baseUrl, nil)
	if err != nil {
		return nil, err
	}

	// add query parameters from map to request
	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := bp.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r, err := unmarshalListResponse(data)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func unmarshalListResponse(result []byte) (*ListResponse, error) {
	d := ListResponse{}
	err := json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
