package equipment

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for equipment
	Endpoint      = "/equipment"
	ChildEndpoint = Endpoint + "/"
)

type (
	EquipmentDetail struct {
		Index             string       `json:"index"`
		Name              string       `json:"name"`
		Url               string       `json:"url"`
		EquipmentCategory api.Ref      `json:"equipment_category"`
		WeaponCategory    string       `json:"weapon_category"`
		WeaponRange       string       `json:"weapon_range"`
		CategoryRange     string       `json:"category_range"`
		Cost              api.Cost     `json:"cost"`
		Damage            DamageDetail `json:"damage"`
		Range             api.Range    `json:"range"`
		Weight            int          `json:"weight"`
		Properties        []api.Ref    `json:"properties"`
	}

	DamageDetail struct {
		DamageDice string
		DamageType api.Ref
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
