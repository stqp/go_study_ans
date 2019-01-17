package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	//local = "http://localhost:9999/gittrend.html"
	url := "https://github.com/trending"
	res, err := http.Get(url)
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

	var repo_names []string
	doc.Find(".repo-list h3 a").Each(func(_ int, s *goquery.Selection) {
		repo_names = append(repo_names, strings.Join(strings.Fields(s.Text()), " "))
	})

	var stars []string
	doc.Find(".repo-list span.float-sm-right").Each(func(_ int, s *goquery.Selection) {
		stars = append(stars, strings.Join(strings.Fields(s.Text()), " "))
	})

	for i := 0; i < len(stars) && i < 10; i++ {
		fmt.Println(repo_names[i] + "　　→　　" + stars[i])
	}

}
