package modules

import (
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
	List string
}

func Bangumilist() []Timeline {
	bilibili := []Timeline{}
	tempacg := Timeline{}
	doc, err := goquery.NewDocument("http://bangumi.bilibili.com/anime/timeline")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".bangumi-list").Each(func(i int, contentSelection *goquery.Selection) {
		c := contentSelection.Find(".side-l")
		tempacg.WeekDay = c.Text()
		tempacg.Url, _ = contentSelection.Find(".preview").Attr("href")
		tempacg.Title, _ = contentSelection.Find(".preview").Attr("title")
		tempacg.Img, _ = contentSelection.Find("img").Attr("src")
		tempacg.num = contentSelection.Find("span").Text()
	})
	bilibili = append(bilibili, tempacg)
	return bilibili
}
