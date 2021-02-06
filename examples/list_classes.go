package main

import (
	"log"
	"net/http"

	go_dnd5eapi "github.com/kjkondratuk/go-dnd5eapi"
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

func listClasses() {
	httpclient := &http.Client{}
	client := go_dnd5eapi.NewClient(api.NewBasicsProvider(httpclient, "https://www.dnd5eapi.co/api"))
	list, err := client.Classes.GetList()
	if err != nil {
		// ... handle error ...
	}
	log.Println("Classes:")
	for _, c := range list.Results {
		log.Println(" - " + c.Name)
	}
}
