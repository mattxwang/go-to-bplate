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
	RecipeLink string
	Location string
	DietaryInfo []string
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

func getItemDietaryInfo(menuItem *goquery.Selection) []string{
	dietInfo := []string{}
	// note: could also use .tooltip-target-wrapper
	menuItem.Find(".item-description").Each(func(i int, itemDescription *goquery.Selection){
		itemDescription.Find("img.webcode-16px").Each(func(j int, dietaryImage *goquery.Selection) {
			infoType, exists := dietaryImage.Attr("alt")
			if exists {
				switch infoType {
				case "V":
					dietInfo = append(dietInfo,"Vegetarian")
				case "VG":
					dietInfo = append(dietInfo,"Vegan")
				case "APNT":
					dietInfo = append(dietInfo,"Peanuts")
				case "ATNT":
					dietInfo = append(dietInfo,"TreeNuts")
				case "AWHT":
					dietInfo = append(dietInfo,"Wheat")
				case "AGTN":
					dietInfo = append(dietInfo,"Gluten")
				case "ASOY":
					dietInfo = append(dietInfo,"Soy")
				case "AMLK":
					dietInfo = append(dietInfo,"Dairy")
				case "AEGG":
					dietInfo = append(dietInfo,"Eggs")
				case "ACSF":
					dietInfo = append(dietInfo,"Shellfish")
				case "AFSH":
					dietInfo = append(dietInfo,"Fish")
				case "HAL":
					dietInfo = append(dietInfo,"Halal")
				case "LC":
					dietInfo = append(dietInfo,"LowCarbonFootprint")
				}
			}
		})
	})
	return dietInfo
}

func getMenuItems(doc *goquery.Document) []MenuItem {
	answers := []MenuItem{}
	doc.Find(".menu-block").Each(func(i int, menuBlock *goquery.Selection){
		location := menuBlock.Find(".col-header").Text()
		if location == "FEAST at Rieber" {
			location = "Feast"
		}
		menuBlock.Find(".menu-item").Each(func(j int, menuItem *goquery.Selection) {
			itemLink := menuItem.Find(".recipelink")
			name := strings.TrimSpace(itemLink.Text())
			recipeLink, exists := itemLink.Attr("href")
			if !exists {
				recipeLink = "#"
			}
			dietInfo := getItemDietaryInfo(menuItem)
			item := MenuItem { 
				Name: name,
				RecipeLink: recipeLink,
				Location: location,
				DietaryInfo: dietInfo,
			}
			answers = append(answers, item)
		})
	})
	return answers
}

func getPageSchedule(doc *goquery.Document) string {
	return doc.Find("#page-header").Text()
}

func filterItemsByKeyword(parents []MenuItem, keywords []string) []MenuItem {
	matches := []MenuItem{}
	for _, parent := range parents {
		parentName := strings.ToLower(parent.Name)
        for _, keyword := range keywords {
			if strings.Contains(parentName, strings.ToLower(keyword)){
				matches = append(matches, parent)
				break
			}
		}
	}
	return matches
}

func filterItemsByDietaryInfo(parents []MenuItem, filters []string) []MenuItem {
	matches := []MenuItem{}
	for _, parent := range parents {
		if len(intersection(parent.DietaryInfo, filters)) > 0 {
			matches = append(matches, parent)
		}
	}
	return matches
}

func xfilterItemsByDietaryInfo(parents []MenuItem, filters []string) []MenuItem {
	matches := []MenuItem{}
	for _, parent := range parents {
		if len(intersection(parent.DietaryInfo, filters)) == 0 {
			matches = append(matches, parent)
		}
	}
	return matches
}

func printMatchesForMeal(date string, meal string, keywords []string, filters []string, xfilters []string) {
	fmt.Println("==========")
	doc := makeHttpRequest("http://menu.dining.ucla.edu/Menus/" + date + "/" + meal)
	title := getPageSchedule(doc)
	matches := getMenuItems(doc)
	if len(keywords) > 0 {
		matches = filterItemsByKeyword(matches, keywords)
	}
	if len(filters) > 0 {
		matches = filterItemsByDietaryInfo(matches, filters)
	}
	if len(xfilters) > 0 {
		matches = xfilterItemsByDietaryInfo(matches, xfilters)
	}
	fmt.Println(title)
	fmt.Println("-------")
	for _, match := range matches {
		fmtString := match.Name + " at " + match.Location 
		for _, dietaryInfo := range match.DietaryInfo {
			fmtString = fmtString + " (" + dietaryInfo + ")"
		}
		fmtString = fmtString + " (" + match.RecipeLink + ")"
		fmt.Println(fmtString)
	}
}
