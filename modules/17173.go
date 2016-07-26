package modules

import (
	"fmt"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

//version 2 直接解析
type Index struct {
	Title  string
	Source string
	Img    string
	Next   string
	Url    string
}
type Page struct {
	PageUp   string
	PageDown string
}

func News_index(url string) ([]Index, []Page) {
	var acg = []Index{}
	var page = []Page{}
	tempacg := Index{}
	temppage := Page{}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".comm-plist3 .item").Each(func(i int, contentSelection *goquery.Selection) {
		c := contentSelection.Find(" .c1 a")
		tempacg.Title = c.Find(".tit").Text()
		tempacg.Url, _ = c.Attr("href")
		tempacg.Img, _ = c.Find("img").Attr("src")
		tempacg.Source = contentSelection.Find(".c2 p").Text()
		acg = append(acg, tempacg)
	})
	doc.Find(".pagination").Each(func(i int, contentSelection *goquery.Selection) {
		c := contentSelection.Find(".prev a")
		s, _ := c.Attr("href")
		r := regexp.MustCompile(`index_(\d)`)
		t := r.FindStringSubmatch(s)
		if len(t) == 0 {
			temppage.PageUp = "begin"
		} else {
			temppage.PageUp = r.FindStringSubmatch(s)[1]
		}
		c = contentSelection.Find(" .next a")
		s, _ = c.Attr("href")
		r = regexp.MustCompile(`index_([0-9]{1,2})`)
		t = r.FindStringSubmatch(s)
		if len(t) == 0 {
			temppage.PageDown = "end"
		} else {
			temppage.PageDown = r.FindStringSubmatch(s)[1]
		}

		page = append(page, temppage)
		fmt.Println(page)
	})

	// fmt.Println(acg)
	return acg, page
}
