// +build integration

// This file is generated by go-dnd5eapi/gen/api_gen.go.  DO NOT EDIT THIS FILE.  Generation parameters are:
// Package Name: monsters
// Endpoint: monsters
// API Name: Monster
// LC API Name: monster

package monsters

import (
    json "github.com/json-iterator/go"
    "fmt"
    "log"
	"net/http"
	"os"
	"testing"

	"github.com/kjkondratuk/go-dnd5eapi/api"
	"github.com/stretchr/testify/assert"
)

var (
	APIBaseURL = os.Getenv("API_ROOT")
	Client     = NewClient(api.NewBasicsProvider(&http.Client{}, APIBaseURL))
)

func TestMain(t *testing.M) {
	t.Run()
}

func TestMonster_GetList_IT(t *testing.T) {
	resp, err := Client.GetList()
	r, _ := json.Marshal(resp)
	log.Print(fmt.Sprintf("Response: %s", r))
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

func TestMonster_GetByIndex_IT(t *testing.T) {
	resp, err := Client.GetByIndex("aboleth")
	r, _ := json.Marshal(resp)
	log.Print(fmt.Sprintf("Response: %s", r))
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}
