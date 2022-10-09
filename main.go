package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
)

var url = "https://zalando.no"

func FetchParse(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	/*
		html, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", html)
	*/

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var uri []string
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		if typee, _ := s.Attr("type"); typee == "text/javascript" {
			textjavascript, _ := s.Attr("src")
			uri = append(uri, textjavascript)
		}
	})

	var fullURL []string
	for _, uri := range uri {
		fullURL := url + uri
		response, err := http.Get(fullURL)
		if err != nil {
			panic(err)
		}

		html, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		if strings.Contains(string(html), "bmak") == true {
			fmt.Println(fullURL)
			break
		} else {
			continue
		}
	}
	return fullURL
}

func main() {
	FetchParse(url)
}
