package main

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetTop100(t *testing.T) {
	setup()
	defer teardown()

	maxItemID := 8435557

	// Set up API stub
	mux.HandleFunc("/v0/maxitem.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, maxItemID)
	})

	// Initialize a user with expected values
	expected, err := client.GetItem(maxItemID)

	if err != nil {
		fmt.Println(err)
	}

	// Test GetTop100
	maxItem := client.GetMaxItem()

	// Checks to make sure request equals expected value
	if !reflect.DeepEqual(maxItem, expected) {
		t.Errorf("client.GetMaxItem()) returned %+v, was expecting %+v", maxItem, expected)
	}
}
