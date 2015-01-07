package gophernews

type Part struct {
	By     string `json:"by"`
	ID     int    `json:"id"`
	Parent int    `json:"parent"`
	Score  int    `json:"score"`
	Text   string `json:"text"`
	Time   int    `json:"time"`
	Type   string `json:"type"`
}
