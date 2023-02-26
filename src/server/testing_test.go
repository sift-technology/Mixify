package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aarti2626/Mixify/src/server/api"
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
		t.Errorf("Incompatible cast from json.Number to int")
	} else {
		fmt.Println("Test succeeded")
	}

}
