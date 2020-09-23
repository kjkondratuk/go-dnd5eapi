package go_dnd5eapi

import "github.com/kjkondratuk/go-dnd5eapi/api"

type (
	ApiClient interface {
		NewClient() ApiClient
		GetList() (*api.ListResponse, error)
		// GetByIndex(index string) (*)
	}

	associationClient struct {
		// ca
	}
	AssociationClient interface {
	}
)
