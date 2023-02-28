package main

import (
	"encoding/json"
	"fmt"
	"strconv"
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

func TestWeightFunct(t *testing.T) {

	var user api.Response

	for i3 := 1; i3 <= 100; i3++ {
		for i1 := 1; i1 <= 4; i1++ {
			for i2 := 1; i2 <= 4; i2++ {
				for i4 := 1; i4 <= 4; i4++ {
					user.R1 = json.Number(strconv.FormatInt(int64(i1), 10))
					user.R2 = json.Number(strconv.FormatInt(int64(i2), 10))
					user.R3 = json.Number(strconv.FormatInt(int64(i3), 10))
					user.R4 = json.Number(strconv.FormatInt(int64(i4), 10))
					api.Weights(&user)

					if user.M.Danceability <= 0 || user.M.Danceability >= 1 || user.M.Energy <= 0 || user.M.Energy >= 1 ||
						user.M.Popularity <= 0 || user.M.Popularity >= 100 || user.M.Acousticness <= 0 || user.M.Acousticness >= 1 {
						t.Errorf("Metric out of range")
					}
				}
			}
		}
	}
	fmt.Println("Test passed: All weights are in range")
}
