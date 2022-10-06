package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

var url string

func fetchParse() {
	fmt.Println("What website would you like to target?")
	fmt.Scanln(&url)
	resp, err := http.Get("https://" + url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		if typee, _ := s.Attr("type"); typee == "text/javascript" {
			textjavascript, _ := s.Attr("src")
			fmt.Printf("https://%s%s\n", url, textjavascript)
		}
	})
}

func main() {
	fetchParse()
}
