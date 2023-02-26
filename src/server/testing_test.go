package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aarti2626/Mixify/src/server/api"
	"github.com/google/uuid"
)

func TestValidateJSONResponse(t *testing.T) {
	var i api.Response
	var user = make(map[string]json.Number)
	user["R1"] = `json:"0"`
	user["R2"] = `json:"1"`
	user["R3"] = `json:"2"`
	user["R4"] = `json:"300"`
	i.R1 = user["R1"]
	i.R2 = user["R2"]
	i.R3 = user["R3"]
	i.R4 = user["R4"]
	if i.R2 != user["R2"] {
		t.Errorf("Could not populate Response with JSON numbers")
	} else {
		fmt.Println("Test passed: valid JSON response")
	}

}

func TestCreateUniqueUUID(t *testing.T) {
	var i api.Response
	i.ID = uuid.New()
	api.NewServer().Responses_DB = append(api.NewServer().Responses_DB, i)
	for index := range api.NewServer().Responses_DB {
		if api.NewServer().Responses_DB[index].ID == i.ID {
			t.Errorf("Non-unique ID found")
		}
	}
	fmt.Println("Test passed: unique UUID created")
}
