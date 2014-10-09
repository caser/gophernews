package main

import (
	// "fmt"
	// "net/http/httptest"
	// "net/http"
	"reflect"
	"testing"
)

func TestGetComment(t *testing.T) {
	client := NewClient()

	expected := Comment{
		By:     "norvig",
		Id:     2921983,
		Kids:   []int{2922097, 2922429, 2924562, 2922709, 2922573, 2922140, 2922141},
		Parent: 2921506,
		Text:   "Aw shucks, guys ... you make me blush with your compliments.<p>Tell you what, Ill make a deal: I'll keep writing if you keep reading. K?",
		Time:   1314211127,
	}

	c, err := client.GetComment(2921983)

	// Makes sure an error wasn't passed
	if err != nil {
		t.Errorf("Error for client.GetComment(2921983) should have been nil. Was: %v", err)
	}

	// Checks to make sure request equals expected value
	if !reflect.DeepEqual(c, expected) {
		t.Errorf("client.GetComment(2921983) returned: \n %+v \nwas expecting: \n %+v", c, expected)
	}

	// Test GetComment with an ID from a non-Story object
	c, err = client.GetComment(8412605)
	// Makes sure an error was passed
	if err == nil {
		t.Errorf("Error for client.GetComment(2921983) should not have been nil. Should have been a type error.")
	}

	// Checks to make sure method returns an empty Story object if the ID is bad
	empty := Comment{}
	if !reflect.DeepEqual(c, empty) {
		t.Errorf("client.GetComment(2921983) returned %+v, should have been empty: %+v", c, empty)
	}
}
