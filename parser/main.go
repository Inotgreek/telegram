package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func check(error error) {
	if error != nil {
		fmt.Println(error)
	}
}

func main() {
	url := "https://www.ebay.com/sch/i.html?_from=R40&_nkw=hd7850&_sacat=0&LH_TitleDesc=0&LH_ItemCondition=3000&_sop=10&rt=nc&LH_BIN=1"
	response := getHtml(url)
	defer response.Body.Close()
	doc, error := goquery.NewDocumentFromReader(response.Body)
	check(error)

	scrapePageData(doc)

}

func getHtml(url string) *http.Response {
	response, error := http.Get(url)
	check(error)

	if response.StatusCode > 400 {
		log.Fatal("Status code:", response.StatusCode)
	}
	return response
}

func scrapePageData(doc *goquery.Document) {
	doc.Find("ul.srp-results>li.s-item").Each(func(index int, item *goquery.Selection) {
		a := item.Find("a.s-item__link")

		title := strings.TrimSpace(a.Text())
		url, _ := a.Attr("href")

		price := item.Find("span.s-item__price").Text()

		scrapedData := []string{title, price, url}
		fmt.Println(scrapedData)

	})

}
