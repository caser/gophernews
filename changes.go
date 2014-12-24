package gophernews

type Changes struct {
	Items    []int    `json:"items"`
	Profiles []string `json:"profiles"`
}
