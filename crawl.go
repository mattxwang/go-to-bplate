package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/PuerkitoBio/goquery"
)

const diningURL = "http://menu.dining.ucla.edu/Menus/"

type MenuItem struct {
	Name string
	RecipeLink string
	Location string
	DietaryInfo []string
}

type MealData struct {
	Title string
	Date string
	Items []MenuItem
}

type DayData struct {
	Date string
	Time time.Time
	Breakfast *MealData
	Lunch *MealData
	Dinner *MealData
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

func filterMenuItems(parents []MenuItem, keywords []string, filters []string, xfilters []string) []MenuItem {
	matches := parents 
	if len(keywords) > 0 {
		matches = filterItemsByKeyword(matches, keywords)
	}
	if len(filters) > 0 {
		matches = filterItemsByDietaryInfo(matches, filters)
	}
	if len(xfilters) > 0 {
		matches = xfilterItemsByDietaryInfo(matches, xfilters)
	}
	return matches
}

func newMealData(title string, date string, items []MenuItem) *MealData {
	m := new(MealData)
	m.Title = title 
	m.Date = date
	m.Items = items
	return m
}

func fetchMealData(date string, meal string, keywords []string, filters []string, xfilters []string) *MealData {
	doc := makeHttpRequest(diningURL + date + "/" + meal)
	title := getPageSchedule(doc)
	items := getMenuItems(doc)
	matches := filterMenuItems(items, keywords, filters, xfilters)
	return newMealData(title, date, matches)
}

func newDayData(date string, breakfast *MealData, lunch *MealData, dinner *MealData) *DayData {
	d := new(DayData)
	d.Date = date 
	d.Time = time.Now()
	d.Breakfast = breakfast 
	d.Lunch = lunch
	d.Dinner = dinner
	return d
}

func fetchDayData(date string, keywords []string, filters []string, xfilters []string) *DayData {
	breakfast := fetchMealData(date, "Breakfast", keywords, filters, xfilters)
	lunch := fetchMealData(date, "Lunch", keywords, filters, xfilters)
	dinner := fetchMealData(date, "Dinner", keywords, filters, xfilters)
	return newDayData(date, breakfast, lunch, dinner)
}

func serializeMealData(mealData *MealData) {
	meal := *mealData
	fmt.Println("==========")
	fmt.Println(meal.Title)
	fmt.Println("-------")
	for _, item := range meal.Items {
		fmtString := item.Name + " at " + item.Location 
		for _, dietaryInfo := range item.DietaryInfo {
			fmtString = fmtString + " (" + dietaryInfo + ")"
		}
		fmtString = fmtString + " (" + item.RecipeLink + ")"
		fmt.Println(fmtString)
	}
}
