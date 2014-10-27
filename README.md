# GopherNews
gophernews is a Go library for accessing the HackerNews API. 

## Usage
```go
import(
  "github.com/caser/gophernews"
)
```

To use gophernews, first construct a new HackerNews API client, then use the various methods to access different parts of the API. 

For example, to see the IDs of all the posts ever submitted by Paul Graham:

```go
client := gophernews.NewClient()
user, err := client.GetUser("pg")
fmt.Println(user.Submitted)
```

In the above example, "pg" is the ID (username) of the user. 

Other accessor methods include:

```go
story, err := client.GetStory(8412605) //=> Returns a Story struct
comment, err := client.GetComment(2921983) //=> Returns a Comment struct
poll, err := client.GetPoll(126809) //=> Returns a Poll struct
part, err := client.GetPart(160705) //=> Returns a Part struct
```

## Special Methods
The HackerNews API also has a few special methods. 

`client.GetTop100()` will return the IDs of the top 100 stories currently trending on Hacker News.

`client.GetMaxItem()` will return the ID of the item (story, comment, etc.) with the largest ID (i.e. the item that was created most recently).

`client.GetChanges()` will return a list of IDs for items and user profiles that were recently changed. The response will be of type:

```go
type Changes struct {
  Items    []int
  Profiles []string
}
```

## Data Structure
---
When you make a Get request for one of the above, a struct will be initialized with the API response. 

Here are the underlying data structues:

```go
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
```

## Next Steps
This is a quick and dirty implementation of the API. Next steps will involve methods for querying relationships (story.Comments, etc.) and more info using the API (rate limiting, etc.). 

Contributions welcome! Write tests, implement feature, send PR.


