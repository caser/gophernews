package gophernews

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetPoll(t *testing.T) {
	setup()
	defer teardown()

	jsonPoll := `{"by":"pg","id":126809,"kids":[126822,126823,126993,126824,126934,127411,126888,127681,126818,126816,126854,127095,126861,127313,127299,126859,126852,126882,126832,127072,127217,126889,127535,126917,126875],"parts":[126810,126811,126812],"score":46,"text":"","time":1204403652,"title":"Poll: What would happen if News.YC had explicit support for polls?","type":"poll"}`

	// Set up API stub
	mux.HandleFunc("/v0/item/126809.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, jsonPoll)
	})

	// Initialize a poll with expected values
	expected := Poll{}

	_ = json.Unmarshal([]byte(jsonPoll), &expected)

	// Test GetPoll with an actual Poll's ID
	p, err := client.GetPoll(126809)

	// Makes sure an error wasn't passed
	if err != nil {
		t.Errorf("Error for client.GetPoll(126809) should have been nil. Was: %v", err)
	}

	// Checks to make sure request equals expected value
	if !reflect.DeepEqual(p, expected) {
		t.Errorf("client.GetPoll(126809) returned %+v, was expecting %+v", p, expected)
	}

	badResponse := `{
    "by" : "dhouston",
    "type" : "comment"
  }`

	mux.HandleFunc("/v0/item/8952.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, badResponse)
	})

	// Test GetPoll with an ID from a non-Poll object
	p, err = client.GetPoll(8952)
	// Makes sure an error was passed
	if err == nil {
		t.Errorf("Error for client.GetPoll(8952) should not have been nil. Should have been a type error.")
	}

	// Checks to make sure method returns an empty Poll object if the ID is bad
	empty := Poll{}
	if !reflect.DeepEqual(p, empty) {
		t.Errorf("client.GetPoll(126809) returned %+v, should have been empty: %+v", p, empty)
	}
}
