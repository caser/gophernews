// TEST CASES

// Test that getStory (comment, part, etc.) return error when ID is of a different type
// and give back the type of the object with the given ID

// Test other cases of wrong input (arguments of wrong type, etc.)

package main

import (
	// "fmt"
	// "net/http/httptest"
	// "net/http"
	"reflect"
	"testing"
)

func TestGetStory(t *testing.T) {
	client := NewClient()

	// Initialize a story with expected values
	expected := Story{
		By: "dhouston",
		Id: 8863,
		Kids: []int{8952, 9224, 8917, 8884, 8887, 8943, 8869,
			8958, 9005, 9671, 8940, 9067, 8908, 9055, 8865, 8881,
			8872, 8873, 8955, 10403, 8903, 8928, 9125, 8998, 8901,
			8902, 8907, 8894, 8878, 8870, 8980, 8934, 8876},
		Score: 111,
		Time:  1175714200,
		Title: "My YC app: Dropbox - Throw away your USB drive",
		Url:   "http://www.getdropbox.com/u/2/screencast.html",
	}

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

/*
  s, err := client.GetStory(8412605) //=> Actual Story
  // s, err := client.GetStory(2921983) //=> Comment (wrong type)
  // c, err := client.GetComment(2921983) //=> Actual Comment
  // c, err := client.GetComment(8412605) //=> Story (wrong type)
  // p, err := client.GetPoll(126809) //=> Actual Poll
  // p, err := client.GetPart(8412605) //=> Story (wrong type)
  // pp, err := client.GetPart(160705) //=> Actual Part of Poll
  // pp, err := client.GetPart(8412605) //=> Story (wrong type)
  // u, err := client.GetUser("pg") //=> User
*/
