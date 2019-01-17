package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {

	res, err := http.Get("https://recruit-tech.co.jp/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		count++
	})
	fmt.Println(count)
}

func main() {
	ExampleScrape()
}
