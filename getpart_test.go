package gophernews

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetPart(t *testing.T) {
	setup()
	defer teardown()

	jsonPart := `{"by":"pg","id":160705,"parent":160704,"score":335,"text":"Yes, ban them; I'm tired of seeing Valleywag stories on News.YC.","time":1207886576,"type":"pollopt"}`

	// Set up API stub
	mux.HandleFunc("/v0/item/160705.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, jsonPart)
	})

	// Initialize a part with expected values
	expected := Part{}

	_ = json.Unmarshal([]byte(jsonPart), &expected)

	// Test GetPart with an actual Part's ID
	p, err := client.GetPart(160705)

	// Makes sure an error wasn't passed
	if err != nil {
		t.Errorf("Error for client.GetPart(160705) should have been nil. Was: %v", err)
	}

	// Checks to make sure request equals expected value
	if !reflect.DeepEqual(p, expected) {
		t.Errorf("client.GetPart(160705) returned %+v, was expecting %+v", p, expected)
	}

	badResponse := `{
    "by" : "dhouston",
    "type" : "comment"
  }`

	mux.HandleFunc("/v0/item/8952.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, badResponse)
	})

	// Test GetPart with an ID from a non-Part object
	p, err = client.GetPart(8952)
	// Makes sure an error was passed
	if err == nil {
		t.Errorf("Error for client.GetPoll(8952) should not have been nil. Should have been a type error.")
	}

	// Checks to make sure method returns an empty Part object if the ID is bad
	empty := Part{}
	if !reflect.DeepEqual(p, empty) {
		t.Errorf("client.GetPoll(126809) returned %+v, should have been empty: %+v", p, empty)
	}
}
