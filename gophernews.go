package gophernews

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "reflect"
	"strconv"
)

// create data structures

type Client struct {
	BaseURI string
	Version string
	Suffix  string
}

type Story struct {
	By    string
	Id    int
	Kids  []int
	Score int
	Time  int
	Title string
	Url   string
}

type Comment struct {
	By     string
	Id     int
	Kids   []int
	Parent int
	Text   string
	Time   int
}

type Poll struct {
	By    string
	Id    int
	Kids  []int
	Parts []int
	Score int
	Text  string
	Time  int
	Title string
}

type Part struct {
	By     string
	Id     int
	Parent int
	Score  int
	Text   string
	Time   int
}

type User struct {
	About     string
	Created   int
	Delay     int
	Id        string
	Karma     int
	Submitted []int
}

type Item struct {
	By      string
	Deleted string
	Id      int
	Kids    []int
	Score   int
	Time    int
	Title   string
	Type    string
	Url     string
	Text    string
	Parent  int
	Parts   []int
}

type Changes struct {
	Items    []int
	Profiles []string
}

// Initializes and returns an API client
func NewClient() Client {
	var c Client
	c.BaseURI = "https://hacker-news.firebaseio.com/"
	c.Version = "v0"
	c.Suffix = ".json"
	return c
}

// Makes an API request and puts response into a Story struct
func (c Client) GetStory(id int) (Story, error) {
	item, err := c.GetItem(id)

	if err != nil {
		return Story{}, err
	}

	if item.Type != "story" {
		emptyStory := Story{}
		return emptyStory, fmt.Errorf("Called GetStory on ID #%v which is not a _story_. "+
			"Item is of type _%v_.", id, item.Type)
	} else {
		story := item.ToStory()
		return story, nil
	}
}

// Makes an API request and puts response into a Comment struct
func (c Client) GetComment(id int) (Comment, error) {
	item, err := c.GetItem(id)

	if err != nil {
		return Comment{}, err
	}

	if item.Type != "comment" {
		emptyComment := Comment{}
		return emptyComment, fmt.Errorf("Called GetComment on ID #%v which is not a _comment_. "+
			"Item is of type _%v_.", id, item.Type)
	} else {
		comment := item.ToComment()
		return comment, nil
	}
}

// Makes an API request and puts response into a Poll struct
func (c Client) GetPoll(id int) (Poll, error) {
	item, err := c.GetItem(id)

	if err != nil {
		return Poll{}, err
	}

	if item.Type != "poll" {
		emptyPoll := Poll{}
		return emptyPoll, fmt.Errorf("Called GetPoll on ID #%v which is not a _poll_. "+
			"Item is of type _%v_.", id, item.Type)
	} else {
		poll := item.ToPoll()
		return poll, nil
	}
}

// Makes an API request and puts response into a Part struct
func (c Client) GetPart(id int) (Part, error) {
	item, err := c.GetItem(id)

	if err != nil {
		return Part{}, err
	}

	if item.Type != "pollopt" {
		emptyPart := Part{}
		return emptyPart, fmt.Errorf("Called GetPart on ID #%v which is not a _part_. "+
			"Item is of type _%v_.", id, item.Type)
	} else {
		part := item.ToPart()
		return part, nil
	}
}

// Makes an API request and puts response into a User struct
func (c Client) GetUser(id string) (User, error) {
	// TODO - refactor URL call into separate method
	url := c.BaseURI + c.Version + "/user/" + id + c.Suffix

	body := c.MakeHTTPRequest(url)

	var u User

	err := json.Unmarshal(body, &u)
	if err != nil {
		fmt.Println(err)
	}

	// TODO - other checking around errors (wrong type, nonexistent user, etc.)
	return u, nil
}

// Makes an API request and puts response into a Item struct
// Items are then converted into Stories, Comments, Polls, and Parts (of polls)
func (c Client) GetItem(id int) (Item, error) {
	url := c.BaseURI + c.Version + "/item/" + strconv.Itoa(id) + c.Suffix

	body := c.MakeHTTPRequest(url)

	var i Item

	if string(body) == "404 page not found" {
		return i, fmt.Errorf("404 page not found")
	}

	err := json.Unmarshal(body, &i)

	if err != nil {
		fmt.Println(err)
	}

	return i, nil
}

func (c Client) GetTop100() ([]int, error) {
	url := c.BaseURI + c.Version + "/topstories" + c.Suffix

	body := c.MakeHTTPRequest(url)

	var top100 []int

	err := json.Unmarshal(body, &top100)

	if err != nil {
		fmt.Println(err)
	}

	return top100, nil
}

func (c Client) GetMaxItem() Item {
	url := c.BaseURI + c.Version + "/maxitem" + c.Suffix

	body := c.MakeHTTPRequest(url)

	var maxItemId int

	err := json.Unmarshal(body, &maxItemId)

	if err != nil {
		fmt.Println(err)
	}

	maxItem, err := c.GetItem(maxItemId)

	if err != nil {
		fmt.Println(err)
	}

	return maxItem
}

func (c Client) GetChanges() (Changes, error) {
	url := c.BaseURI + c.Version + "/updates" + c.Suffix

	body := c.MakeHTTPRequest(url)

	var changes Changes

	err := json.Unmarshal(body, &changes)

	if err != nil {
		fmt.Println(err)
	}

	return changes, nil
}

func (c Client) MakeHTTPRequest(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}

// Convert an item to a Story
func (i Item) ToStory() Story {
	var s Story
	s.By = i.By
	s.Id = i.Id
	s.Kids = i.Kids
	s.Score = i.Score
	s.Time = i.Time
	s.Title = i.Title
	s.Url = i.Url
	return s
}

// Convert an item to a Comment
func (i Item) ToComment() Comment {
	var c Comment
	c.By = i.By
	c.Id = i.Id
	c.Kids = i.Kids
	c.Parent = i.Parent
	c.Text = i.Text
	c.Time = i.Time
	return c
}

// Convert an item to a Poll
func (i Item) ToPoll() Poll {
	var p Poll
	p.By = i.By
	p.Id = i.Id
	p.Kids = i.Kids
	p.Parts = i.Parts
	p.Score = i.Score
	p.Text = i.Text
	p.Time = i.Time
	p.Title = i.Title
	return p
}

// Convert an item to a Part
func (i Item) ToPart() Part {
	var p Part
	p.By = i.By
	p.Id = i.Id
	p.Parent = i.Parent
	p.Score = i.Score
	p.Text = i.Text
	p.Time = i.Time
	return p
}

func main() {
	client := NewClient()

	s, err := client.GetStory(8412605) //=> Actual Story
	// s, err := client.GetStory(2921983) //=> Comment (wrong type)
	// c, err := client.GetComment(2921983) //=> Actual Comment
	// c, err := client.GetComment(8412605) //=> Story (wrong type)
	// p, err := client.GetPoll(126809) //=> Actual Poll
	// p, err := client.GetPart(8412605) //=> Story (wrong type)
	// pp, err := client.GetPart(160705) //=> Actual Part of Poll
	// pp, err := client.GetPart(8412605) //=> Story (wrong type)
	// u, err := client.GetUser("pg") //=> User

	if err != nil {
		fmt.Println(err)
	} else {
		// fmt.Println(u.About, "\n", u.Created, "\n", u.Karma)
		fmt.Println(s.By, "\n", s.Title, "\n", s.Score)
	}

	// write accessors to get stories, comments, polls, parts, and users
	// write accessors for special cases (top stories, updates, etc.)
	// write special accessors for stories, comments, etc. to get objects instead of
	// IDs (ints) of parents, children, etc.
}
