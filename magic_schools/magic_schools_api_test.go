// +build integration

// This file is generated by go-dnd5eapi/gen/api_gen.go.  DO NOT EDIT THIS FILE.  Generation parameters are:
// Package Name: magic_schools
// Endpoint: magic-schools
// API Name: MagicSchool
// LC API Name: magicSchool

package magic_schools

import (
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

func TestMagicSchool_GetList_IT(t *testing.T) {
	_, err := Client.GetList()
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}

func TestMagicSchool_GetByIndex_IT(t *testing.T) {
	_, err := Client.GetByIndex("abjuration")
	assert.Nil(t, err, "Should not receive an error contacting API.")
	assert.True(t, true, "Should complete successfully!")
}