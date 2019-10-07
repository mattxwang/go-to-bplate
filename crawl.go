package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

type Document struct {}

func makeHttpRequest(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func getMenuItems(doc *goquery.Document) []string {
	answers := []string{}
	doc.Find(".menu-item").Each(func(i int, s *goquery.Selection) {
		item := s.Find(".recipelink")
		// linkLoc, _ := item.Attr("href")
		//if exists {
			// fmt.Println(linkLoc)
			answers = append(answers, strings.TrimSpace(item.Text()))
		//}
	})
	return answers
}

func getMatchingItems(parent []string, keywords []string) []string {
	matches := []string{}
	for _, item := range parent {
        for _, keyword := range keywords {
			if strings.Contains(strings.ToLower(item), strings.ToLower(keyword)){
				matches = append(matches, item)
			}
		}
	}
	return matches
}

func main() {
	keywords := []string{"chicken"}
	doc := makeHttpRequest("http://menu.dining.ucla.edu/Menus")
	items := getMenuItems(doc)
	matches := getMatchingItems(items, keywords)
	fmt.Println(strings.Join(matches, ", "))
}
