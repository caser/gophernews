package gophernews

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetChanges(t *testing.T) {
	setup()
	defer teardown()

	jsonChanges := `{"items":[8435060,8435585,8429118,8435548,8434512,8432703,8427757,8434023,8435467,8435278],"profiles":["vineet","walterbell","briandear","libovness","nfm","integraton","csdrane","tdicola","philipDS","mkremins","hxrts","adventured","kqr2","Renaud"]}`

	// Set up API stub
	mux.HandleFunc("/v0/updates.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"items":[8435060,8435585,8429118,8435548,8434512,8432703,8427757,8434023,8435467,8435278],"profiles":["vineet","walterbell","briandear","libovness","nfm","integraton","csdrane","tdicola","philipDS","mkremins","hxrts","adventured","kqr2","Renaud"]}`)
	})

	// Initialize Changes with expected values

	var expected Changes

	err := json.Unmarshal([]byte(jsonChanges), &expected)
	if err != nil {
		fmt.Println(err)
	}

	// Test GetChanges
	c, err := client.GetChanges()
	if err != nil {
		fmt.Println(err)
	}

	// Checks to make sure request equals expected value
	if !reflect.DeepEqual(c, expected) {
		t.Errorf("client.GetChanges() returned %+v, was expecting %+v", c, expected)
	}
}
