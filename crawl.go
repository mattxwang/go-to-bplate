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
		answers = append(answers,item.Text())
	})
	return answers
}

func main() {
	// keywords := []string{"chicken"}
	doc := makeHttpRequest("http://menu.dining.ucla.edu/Menus")
	items := getMenuItems(doc)
	fmt.Println(strings.Join(items, ", "))
}
