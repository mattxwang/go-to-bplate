package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

type MenuItem struct {
	Name string
	Location string
}

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

func getMenuItems(doc *goquery.Document) []MenuItem {
	answers := []MenuItem{}
	doc.Find(".menu-block").Each(func(i int, block *goquery.Selection){
		location := block.Find(".col-header").Text()
		if location == "FEAST at Rieber" {
			location = "Feast"
		}
		block.Find(".menu-item").Each(func(j int, link *goquery.Selection) {
			name := strings.TrimSpace(link.Find(".recipelink").Text())
			item := MenuItem { 
				Name: name,
				Location: location,
			}
			answers = append(answers, item)
		})
	})
	return answers
}

func getPageSchedule(doc *goquery.Document) string {
	return doc.Find("#page-header").Text()
}

func getMatchingItems(parents []MenuItem, keywords []string) []MenuItem {
	matches := []MenuItem{}
	for _, parent := range parents {
		parentName := strings.ToLower(parent.Name)
        for _, keyword := range keywords {
			if strings.Contains(parentName, strings.ToLower(keyword)){
				matches = append(matches, parent)
			}
		}
	}
	return matches
}

func printMatchesForMeal(meal string, keywords []string){
	fmt.Println("==========")
	doc := makeHttpRequest("http://menu.dining.ucla.edu/Menus/" + meal)
	title := getPageSchedule(doc)
	items := getMenuItems(doc)
	matches := getMatchingItems(items, keywords)
	fmt.Println(title)
	fmt.Println("-------")
	for _, match := range matches {
		fmt.Println(match.Name + " at " + match.Location)
	}
}

func main() {
	keywords := []string{"chicken"}
	printMatchesForMeal("Breakfast", keywords)
	printMatchesForMeal("Lunch", keywords)
	printMatchesForMeal("Dinner", keywords)
}
