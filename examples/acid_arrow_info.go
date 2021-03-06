package main

import (
	"encoding/json"
	"log"
	"net/http"

	go_dnd5eapi "github.com/kjkondratuk/go-dnd5eapi"
	"github.com/kjkondratuk/go-dnd5eapi/api"
)

func acidArrowInfo() {
	httpclient := &http.Client{}
	client := go_dnd5eapi.NewClient(api.NewBasicsProvider(httpclient, "https://www.dnd5eapi.co/api"))
	spellDetail, err := client.Spells.GetByIndex("acid-arrow")
	if err != nil {
		// ... handle error ...
	}
	log.Printf("Name: %s\n", spellDetail.Name)
	str, _ := json.MarshalIndent(spellDetail, "", "  ")
	log.Printf("Raw: %s\n", str)
}
