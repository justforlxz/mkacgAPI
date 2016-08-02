package modules

import (
	"encoding/json"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Timeline struct {
	Title       string
	Url         string
	Img         string
	WeekDay     string
	num         string
	Bangumilist []BangumiList
}
type BangumiList struct {
	List  string
	Title string
}

type Bilibili struct {
	Bangumi bangumi `json:"bangumi"`
}
type bangumi struct {
	List []list
}
type list struct {
	Aera       string `json:"area"`
	Cover      string `json:"cover"`
	Title      string `json:"title"`
	Season_id  int    `json:"season_id"`
	Weekday    int    `json:"weekday"`
	Play_count int    `json:"play_count"`
}

func Bangumilist() Bilibili {
	doc, err := goquery.NewDocument("http://www.bilibili.com/index/index-bangumi-timeline.json")
	if err != nil {
		log.Fatal(err)
	}
	var s Bilibili
	json.Unmarshal([]byte(doc.Text()), &s)
	return s
}
