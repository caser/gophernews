package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetStory(t *testing.T) {
	setup()
	defer teardown()

	// Initialize a story with expected values
	jsonStory := `{"by":"dhouston","id":8863,"kids":[8952,9224,8917,8884,8887,8943,8869,8958,9005,9671,8940,9067,8908,9055,8865,8881,8872,8873,8955,10403,8903,8928,9125,8998,8901,8902,8907,8894,8878,8870,8980,8934,8876],"score":111,"time":1175714200,"title":"My YC app: Dropbox - Throw away your USB drive","type":"story","url":"http://www.getdropbox.com/u/2/screencast.html"}`

	// Set up API stub
	mux.HandleFunc("/v0/item/8863.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, jsonStory)
	})

	// Initialize a story with expected values
	expected := Story{}

	_ = json.Unmarshal([]byte(jsonStory), &expected)

	// Test GetStory with an actual Story's ID
	s, err := client.GetStory(8863)

	// Makes sure an error wasn't passed
	if err != nil {
		t.Errorf("Error for client.GetStory(8863) should have been nil. Was: %v", err)
	}

	// Checks to make sure request equals expected value
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("client.GetStory(8863) returned %+v, was expecting %+v", s, expected)
	}

	badResponse := `{
		"by" : "dhouston",
		"type" : "comment"
	}`

	mux.HandleFunc("/v0/item/8952.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, badResponse)
	})

	// Test GetStory with an ID from a non-Story object
	s, err = client.GetStory(8952)
	// Makes sure an error was passed
	if err == nil {
		t.Errorf("Error for client.GetStory(8952) should not have been nil. Should have been a type error.")
	}

	// Checks to make sure method returns an empty Story object if the ID is bad
	empty := Story{}
	if !reflect.DeepEqual(s, empty) {
		t.Errorf("client.GetStory(8952) returned %+v, should have been empty: %+v", s, empty)
	}
}
