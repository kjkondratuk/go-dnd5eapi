package main

import (
	"log"
	"net/http"

	go_dnd5eapi "github.com/kjkondratuk/go-dnd5eapi"
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

func queryList() {
	httpclient := &http.Client{}
	client := go_dnd5eapi.NewClient(api.NewBasicsProvider(httpclient, "https://www.dnd5eapi.co/api"))
	query := make(map[string]string, 1)
	query["name"] = "Barbarian"
	list, err := client.Classes.QueryList(query)
	if err != nil {
		// ... handle error ...
	}
	log.Println("Queried Classes:")
	for _, c := range list.Results {
		log.Println(" - " + c.Name)
	}
}
