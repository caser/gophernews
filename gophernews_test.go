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
		t.Error(err)
	}

	// Test GetChanges
	c, err := client.GetChanges()
	if err != nil {
		t.Error(err)
	}

	// Checks to make sure request equals expected value
	if !reflect.DeepEqual(c, expected) {
		t.Errorf("client.GetChanges() returned %+v, was expecting %+v", c, expected)
	}
}

func TestGetComment(t *testing.T) {
	setup()
	defer teardown()

	jsonComment := `{"by":"norvig","id":2921983,"kids":[2922097,2922429,2924562,2922709,2922573,2922140,2922141],"parent":2921506,"text":"Aw shucks, guys ... you make me blush with your compliments.<p>Tell you what, Ill make a deal: I'll keep writing if you keep reading. K?","time":1314211127,"type":"comment"}`

	// Set up API stub
	mux.HandleFunc("/v0/item/2921983.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, jsonComment)
	})

	// Initialize a comment with expected values
	expected := Comment{}

	_ = json.Unmarshal([]byte(jsonComment), &expected)

	c, err := client.GetComment(2921983)

	// Makes sure an error wasn't passed
	if err != nil {
		t.Errorf("Error for client.GetComment(2921983) should have been nil. Was: %v", err)
	}

	// Checks to make sure request equals expected value
	if !reflect.DeepEqual(c, expected) {
		t.Errorf("client.GetComment(2921983) returned: \n %+v \nwas expecting: \n %+v", c, expected)
	}

	badResponse := `{
    "by" : "dhouston",
    "type" : "story"
  }`

	// Stup bad API response
	mux.HandleFunc("/v0/item/8412605.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, badResponse)
	})

	// Test GetComment with an ID from a non-Story object
	c, err = client.GetComment(8412605)
	// Makes sure an error was passed
	if err == nil {
		t.Errorf("Error for client.GetComment(8412605) should not have been nil. Should have been a type error.")
	}

	// Checks to make sure method returns an empty Story object if the ID is bad
	empty := Comment{}
	if !reflect.DeepEqual(c, empty) {
		t.Errorf("client.GetComment(8412605) returned %+v, should have been empty: %+v", c, empty)
	}
}

func TestGetMax(t *testing.T) {
	setup()
	defer teardown()

	maxItemID := 8435557

	// Set up API stub
	mux.HandleFunc("/v0/maxitem.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, maxItemID)
	})

	// Set up API stub
	mux.HandleFunc("/v0/item/8435557.json", func(w http.ResponseWriter, r *http.Request) {
		result := `{"by":"jaguar86","id":8435557,"kids":[8435840,8435571,8435665],"parent":8435467,"text":"And they would like to nominate Oracle, IBM and Microsoft to take this challenge within the next 24 hours ...","time":1412898130,"type":"comment"}`
		fmt.Fprint(w, result)
	})

	// Initialize Max Item ID with expected values
	expected, err := client.GetItem(maxItemID)

	if err != nil {
		t.Errorf("Error on %d: %s", maxItemID, err)
	}

	// Test GetMaxItem
	maxItem, err := client.GetMaxItem()

	if err != nil {
		t.Error(err)
	}

	// Checks to make sure request equals expected value
	if !reflect.DeepEqual(maxItem, expected) {
		t.Errorf("client.GetMaxItem()) returned %+v, was expecting %+v", maxItem, expected)
	}
}

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

func TestGetTop100(t *testing.T) {
	setup()
	defer teardown()

	jsonTop100 := `[8434128,8432709,8434649,8432703,8431936,8435278,8434338,8434997,8434785,8434409,8433734,8432423,8431653,8433690,8434268,8433824,8434430,8432373,8428632,8434391,8433505,8432211,8434512,8434966,8432919,8433691,8432838,8432857,8432528,8430412,8435195,8422599,8434996,8433163,8428522,8433046,8431635,8430544,8430096,8433247,8434309,8432305,8434735,8431753,8427852,8427757,8433869,8433816,8435047,8432072,8434403,8422087,8435402,8428056,8432157,8432555,8426148,8434023,8431361,8428849,8434823,8431640,8424696,8431961,8429123,8435307,8430611,8435449,8433681,8432616,8432388,8434830,8422581,8432145,8435168,8434618,8434173,8433703,8434097,8435202,8433879,8435299,8431697,8425501,8434589,8432021,8434350,8434641,8435089,8435158,8434900,8433552,8434905,8434868,8426984,8427086,8422928,8432130,8434950,8431577]`

	// Set up API stub
	mux.HandleFunc("/v0/topstories.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, jsonTop100)
	})

	// Initialize a user with expected values
	expected := []int{8434128, 8432709, 8434649, 8432703, 8431936, 8435278, 8434338, 8434997, 8434785, 8434409, 8433734, 8432423, 8431653, 8433690, 8434268, 8433824, 8434430, 8432373, 8428632, 8434391, 8433505, 8432211, 8434512, 8434966, 8432919, 8433691, 8432838, 8432857, 8432528, 8430412, 8435195, 8422599, 8434996, 8433163, 8428522, 8433046, 8431635, 8430544, 8430096, 8433247, 8434309, 8432305, 8434735, 8431753, 8427852, 8427757, 8433869, 8433816, 8435047, 8432072, 8434403, 8422087, 8435402, 8428056, 8432157, 8432555, 8426148, 8434023, 8431361, 8428849, 8434823, 8431640, 8424696, 8431961, 8429123, 8435307, 8430611, 8435449, 8433681, 8432616, 8432388, 8434830, 8422581, 8432145, 8435168, 8434618, 8434173, 8433703, 8434097, 8435202, 8433879, 8435299, 8431697, 8425501, 8434589, 8432021, 8434350, 8434641, 8435089, 8435158, 8434900, 8433552, 8434905, 8434868, 8426984, 8427086, 8422928, 8432130, 8434950, 8431577}

	// Test GetTop100
	top, err := client.GetTop100()

	// Makes sure an error wasn't passed
	if err != nil {
		t.Errorf("Error when calling GetTop100:\n", err)
	}

	// Checks to make sure request equals expected value
	if !reflect.DeepEqual(top, expected) {
		t.Errorf("client.GetTop100() returned %+v, was expecting %+v", top, expected)
	}
}
