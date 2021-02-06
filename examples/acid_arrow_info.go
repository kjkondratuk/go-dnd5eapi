package main

import (
	"log"
	"net/http"

	go_dnd5eapi "github.com/kjkondratuk/go-dnd5eapi"
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

func acidArrowInfo() {
	httpclient := &http.Client{}
	client := go_dnd5eapi.NewClient(api.NewBasicsProvider(httpclient, "https://dnd5eapi.co/api"))
	spellDetail, err := client.Spells.GetByIndex("acid-arrow")
	if err != nil {
		// ... handle error ...
	}
	log.Printf("Name: %s\n", spellDetail.Name)
	log.Printf("Raw: %+v\n", spellDetail)
}
