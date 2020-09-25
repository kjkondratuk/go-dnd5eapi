package equipment

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for conditions
	Endpoint      = "/equipment"
	ChildEndpoint = Endpoint + "/"
)

type (
	EquipmentDetail struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Url   string `json:"url"`
	}

	equipmentClient struct {
		basicsProvider api.BasicsProvider
	}

	EquipmentClient interface {
		GetList() (*api.ListResponse, error)
		GetByIndex(index string) (*EquipmentDetail, error)
	}
)

func NewClient(basicsProvider api.BasicsProvider) EquipmentClient {
	return &equipmentClient{
		basicsProvider: basicsProvider,
	}
}

func (ac *equipmentClient) GetList() (*api.ListResponse, error) {
	return ac.basicsProvider.GetListForUrl(Endpoint)
}

func (ac *equipmentClient) GetByIndex(index string) (*EquipmentDetail, error) {
	result, err := ac.basicsProvider.ApiGet(ChildEndpoint + index)
	if err != nil {
		return nil, err
	}

	d := EquipmentDetail{}
	err = json.Unmarshal(result, &d)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
