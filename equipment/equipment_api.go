// This file is generated by go-dnd5eapi/gen/api_gen.go.  DO NOT EDIT THIS FILE.  Generation parameters are:
// Package Name: equipment
// Endpoint: equipment
// API Name: Equipment
// LC API Name: equipment
// API Class: EquipmentDetail

package equipment

import (
	"encoding/json"

	"github.com/kjkondratuk/go-dnd5eapi/api"
)

const (
	// Endpoints for ability scores
	Endpoint      = "/equipment"
	ChildEndpoint = Endpoint + "/"
)

type (
	equipmentClient struct {
		basicsProvider api.BasicsProvider
	}
)

type EquipmentClient interface {
	GetList() (*api.ListResponse, error)
	GetByIndex(index string) (*EquipmentDetail, error)
}

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
