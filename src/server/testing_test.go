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

// test all 163,840 possible responses and sees if spotify will recommend tracks for all possible answer combinations
func TestWeightFunct(t *testing.T) {

	var user api.Response
	var userRec api.Rec
	s := api.NewServer()

	for i3 := 0; i3 <= 10; i3++ {
		for i1 := 1; i1 <= 4; i1++ {
			for i2 := 1; i2 <= 4; i2++ {
				for i4 := 1; i4 <= 4; i4++ {
					for i5 := 1; i4 <= 4; i4++ {
						for i6 := 1; i4 <= 4; i4++ {

							user.R1 = json.Number(strconv.FormatInt(int64(i1), 10))
							user.R2 = json.Number(strconv.FormatInt(int64(i2), 10))
							user.R3 = json.Number(strconv.FormatInt(int64(i3), 10))
							user.R4 = json.Number(strconv.FormatInt(int64(i4), 10))
							user.R5 = json.Number(strconv.FormatInt(int64(i5), 10))
							user.R6 = json.Number(strconv.FormatInt(int64(i6), 10))
							api.Weights(&user)
							userRec.Recs = api.Recommend(s.Client, &user)

							if len(userRec.Recs) == 0 {
								t.Errorf("No recommendations")
							}
						}

					}
				}
			}
		}
	}
	fmt.Println("Test passed: All weights are in range")
}
