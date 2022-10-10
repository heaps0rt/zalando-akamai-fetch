package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strings"
)

var url = "ikea.com"

func FetchParse(url string) {
	resp, err := http.Get("https://" + url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// checks html for tag <script type="text/javascript" src=
	var uri []string
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		if typee, _ := s.Attr("type"); typee == "text/javascript" {
			textjavascript, _ := s.Attr("src")
			uri = append(uri, textjavascript)
		}
	})

	for _, uri := range uri {
		fullURL := uri
		if strings.Contains(uri, "https://") == false {
			response, err := http.Get("https://" + url + fullURL)
			if err != nil {
				panic(err)
			}

			html, err := io.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}

			if strings.Contains(string(html), "bmak") == true {
				fmt.Println("https://" + url + fullURL)
				break
			} else {
				continue
			}
		}
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

}

func main() {
	FetchParse(url)

}
