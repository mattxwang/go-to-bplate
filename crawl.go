package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

type DietaryInformation struct {
	Vegetarian bool
	Vegan bool
	Peanuts bool 
	TreeNuts bool 
	Wheat bool 
	Gluten bool 
	Soy bool
	Dairy bool
	Eggs bool
	Shellfish bool
	Fish bool
	Halal bool
	LowCarbon bool
}

type MenuItem struct {
	Name string
	RecipeLink string
	Location string
	DietaryInfo DietaryInformation
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

func getItemDietaryInfo(menuItem *goquery.Selection) DietaryInformation{
	dietInfo := DietaryInformation {
		Vegetarian: false,
		Vegan: false,
		Peanuts: false,
		TreeNuts: false,
		Wheat: false,
		Gluten: false,
		Soy: false,
		Dairy: false,
		Eggs: false,
		Shellfish: false,
		Fish: false,
		Halal: false,
		LowCarbon: false,
	}
	// note: could also use .tooltip-target-wrapper
	menuItem.Find(".item-description").Each(func(i int, itemDescription *goquery.Selection){
		itemDescription.Find("img.webcode-16px").Each(func(j int, dietaryImage *goquery.Selection) {
			infoType, exists := dietaryImage.Attr("alt")
			if exists {
				switch infoType {
				case "V":
					dietInfo.Vegetarian = true 
				case "VG":
					dietInfo.Vegan = true 
				case "APNT":
					dietInfo.Peanuts = true 
				case "ATNT":
					dietInfo.TreeNuts = true 
				case "AWHT":
					dietInfo.Wheat = true 
				case "AGTN":
					dietInfo.Gluten = true 
				case "ASOY":
					dietInfo.Soy = true 
				case "AMLK":
					dietInfo.Dairy = true
				case "AEGG":
					dietInfo.Eggs = true
				case "ACSF":
					dietInfo.Shellfish = true 
				case "AFSH":
					dietInfo.Fish = true
				case "HAL":
					dietInfo.Halal = true
				case "LC":
					dietInfo.LowCarbon = true
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

func searchItemsByKeyword(parents []MenuItem, keywords []string) []MenuItem {
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

func printMatchesForMeal(date string, meal string, keywords []string){
	fmt.Println("==========")
	doc := makeHttpRequest("http://menu.dining.ucla.edu/Menus/" + date + "/" + meal)
	title := getPageSchedule(doc)
	items := getMenuItems(doc)
	matches := searchItemsByKeyword(items, keywords)
	fmt.Println(title)
	fmt.Println("-------")
	for _, match := range matches {
		fmtString := match.Name + " at " + match.Location 
		if match.DietaryInfo.Vegetarian {
			fmtString = fmtString + " (VG)"
		}
		if match.DietaryInfo.Vegan {
			fmtString = fmtString + " (V)"
		}
		fmtString = fmtString + " (" + match.RecipeLink + ")"
		fmt.Println(fmtString)
	}
}

func main() {
	keywords := []string{"chicken", "tacos", "avocado"}
	printMatchesForMeal("Today", "Breakfast", keywords)
	printMatchesForMeal("Today","Lunch", keywords)
	printMatchesForMeal("Today","Dinner", keywords)
}
