package gophernews

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// create data structures

type Client struct {
	BaseURI string
	Version string
	Suffix  string
}

// All the struct definitions can be generated automatically using the example JSON provided by the actual API endpoints corresponding to the test cases
// gojson can be installed with `go get github.com/ChimeraCoder/gojson`

//go:generate gojson -o user.go -name "User" -pkg "gophernews" -input json/chimeracoder.json

//go:generate gojson -o changes.go -name "Changes" -pkg "gophernews" -input json/updates.json

//go:generate gojson -o comment.go -name "Comment" -pkg "gophernews" -input json/2921983.json

//go:generate gojson -o poll.go -name "Poll" -pkg "gophernews" -input json/126809.json

//go:generate gojson -o story.go -name "Story" -pkg "gophernews" -input json/8863.json

//go:generate gojson -o part.go -name "Part" -pkg "gophernews" -input json/160705.json

// Initializes and returns an API client
func NewClient() *Client {
	var c Client
	c.BaseURI = "https://hacker-news.firebaseio.com/"
	c.Version = "v0"
	c.Suffix = ".json"
	return &c
}

// Makes an API request and puts response into a Story struct
func (c *Client) GetStory(id int) (Story, error) {
	item, err := c.GetItem(id)

	if err != nil {
		return Story{}, err
	}

	if item.Type() != "story" {
		emptyStory := Story{}
		return emptyStory, fmt.Errorf("Called GetStory on ID #%v which is not a _story_. "+
			"Item is of type _%v_.", id, item.Type)
	} else {
		story := item.ToStory()
		return story, nil
	}
}

// Makes an API request and puts response into a Comment struct
func (c *Client) GetComment(id int) (Comment, error) {
	item, err := c.GetItem(id)

	if err != nil {
		return Comment{}, err
	}

	if item.Type() != "comment" {
		emptyComment := Comment{}
		return emptyComment, fmt.Errorf("Called GetComment on ID #%v which is not a _comment_. "+
			"Item is of type _%v_.", id, item.Type)
	} else {
		comment := item.ToComment()
		return comment, nil
	}
}

// Makes an API request and puts response into a Poll struct
func (c *Client) GetPoll(id int) (Poll, error) {
	item, err := c.GetItem(id)

	if err != nil {
		return Poll{}, err
	}

	if item.Type() != "poll" {
		emptyPoll := Poll{}
		return emptyPoll, fmt.Errorf("Called GetPoll on ID #%v which is not a _poll_. "+
			"Item is of type _%v_.", id, item.Type)
	} else {
		poll := item.ToPoll()
		return poll, nil
	}
}

// Makes an API request and puts response into a Part struct
func (c *Client) GetPart(id int) (Part, error) {
	item, err := c.GetItem(id)

	if err != nil {
		return Part{}, err
	}

	if item.Type() != "pollopt" {
		emptyPart := Part{}
		return emptyPart, fmt.Errorf("Called GetPart on ID #%v which is not a _part_. "+
			"Item is of type _%v_.", id, item.Type)
	} else {
		part := item.ToPart()
		return part, nil
	}
}

// Makes an API request and puts response into a User struct
func (c *Client) GetUser(id string) (User, error) {
	// TODO - refactor URL call into separate method
	url := c.BaseURI + c.Version + "/user/" + id + c.Suffix

	var u User

	body, err := c.MakeHTTPRequest(url)
	if err != nil {
		return u, err
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		return u, err
	}

	// TODO - other checking around errors (wrong type, nonexistent user, etc.)
	return u, nil
}

// Makes an API request and puts response into a item struct
// items are then converted into Stories, Comments, Polls, and Parts (of polls)
func (c *Client) GetItem(id int) (item, error) {
	url := c.BaseURI + c.Version + "/item/" + strconv.Itoa(id) + c.Suffix

	var i map[string]interface{}

	body, err := c.MakeHTTPRequest(url)
	if err != nil {
		return i, err
	}

	if string(body) == "404 page not found" {
		return i, fmt.Errorf("404 page not found")
	}

	err = json.Unmarshal(body, &i)

	return i, err
}

func (c *Client) GetTop100() ([]int, error) {
	url := c.BaseURI + c.Version + "/topstories" + c.Suffix

	body, err := c.MakeHTTPRequest(url)

	var top100 []int

	err = json.Unmarshal(body, &top100)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return top100, nil
}

func (c *Client) GetMaxItem() (Item, error) {
	url := c.BaseURI + c.Version + "/maxitem" + c.Suffix

	body, err := c.MakeHTTPRequest(url)

	var maxItemId int

	err = json.Unmarshal(body, &maxItemId)
	if err != nil {
		return item{}, err
	}

	maxItem, err := c.GetItem(maxItemId)

	return maxItem, err
}

func (c *Client) GetChanges() (Changes, error) {
	url := c.BaseURI + c.Version + "/updates" + c.Suffix

	body, err := c.MakeHTTPRequest(url)

	var changes Changes

	err = json.Unmarshal(body, &changes)

	return changes, err
}

func (c *Client) MakeHTTPRequest(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf(http.StatusText(http.StatusNotFound))
	}
	return body, nil
}

// Convert an item to a Story
func (i item) ToStory() Story {
	var s Story
	s.By = i.By()
	s.ID = i.ID()
	s.Kids = i.Kids()
	s.Score = i.Score()
	s.Time = i.Time()
	s.Title = i.Title()
	s.Type = i.Type()
	s.URL = i.URL()
	return s
}

// Convert an item to a Comment
func (i item) ToComment() Comment {
	var c Comment
	c.By = i.By()
	c.ID = i.ID()
	c.Kids = i.Kids()
	c.Parent = i.Parent()
	c.Text = i.Text()
	c.Time = i.Time()
	c.Type = i.Type()
	return c
}

// Convert an item to a Poll
func (i item) ToPoll() Poll {
	var p Poll
	p.By = i.By()
	p.ID = i.ID()
	p.Kids = i.Kids()
	p.Parts = i.Parts()
	p.Score = i.Score()
	p.Text = i.Text()
	p.Time = i.Time()
	p.Title = i.Title()
	p.Type = i.Type()
	return p
}

// Convert an item to a Part
func (i item) ToPart() Part {
	var p Part
	p.By = i.By()
	p.ID = i.ID()
	p.Parent = i.Parent()
	p.Score = i.Score()
	p.Text = i.Text()
	p.Time = i.Time()
	p.Type = i.Type()
	return p
}

func main() {
	client := NewClient()

	// README
	s, err := client.GetStory(8412605) //=> Actual Story
	// c, err := client.GetComment(2921983) //=> Actual Comment
	// p, err := client.GetPoll(126809) //=> Actual Poll
	// pp, err := client.GetPart(160705) //=> Actual Part of Poll
	// u, err := client.GetUser("pg") //=> User

	if err != nil {
		panic(err)
	} else {
		// fmt.Println(u.About, "\n", u.Created, "\n", u.Karma)
		fmt.Println(s.By, "\n", s.Title, "\n", s.Score)
	}

	// write accessors to get stories, comments, polls, parts, and users
	// write accessors for special cases (top stories, updates, etc.)
	// write special accessors for stories, comments, etc. to get objects instead of
	// IDs (ints) of parents, children, etc.
}
