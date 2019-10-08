package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	serverModePtr := flag.Bool("s", false, "runs in API webserver mode if true")
	datePtr := flag.String("date", "Today", "the day to poll menu data from")
	keywordsPtr := flag.String("keywords", "", "comma-delimited keywords to search by")
	filtersPtr := flag.String("filters", "", "comma-delimited (inclusive) filters for dietary information")
	xfiltersPtr := flag.String("xfilters", "", "comma-delimited (exclusive) filters for dietary information")
	flag.Parse()
	if *serverModePtr {
		fmt.Println("Server mode")
	} else {
		keywords := strings.Split(*keywordsPtr, ",")
		filters := strings.Split(*filtersPtr, ",")
		if strings.TrimSpace(*filtersPtr) == "" {
			filters = []string{}
		}
		xfilters := strings.Split(*xfiltersPtr, ",")
		if strings.TrimSpace(*xfiltersPtr) == "" {
			xfilters = []string{}
		}
		printMatchesForMeal(*datePtr, "Breakfast", keywords, filters, xfilters)
		printMatchesForMeal(*datePtr,"Lunch", keywords, filters, xfilters)
		printMatchesForMeal(*datePtr,"Dinner", keywords, filters, xfilters)
	}
}