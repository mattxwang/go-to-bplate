package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const portNum = "4242"

func main() {
	serverModePtr := flag.Bool("s", false, "runs in API webserver mode if true")
	datePtr := flag.String("date", "Today", "the day to poll menu data from")
	keywordsPtr := flag.String("keywords", "", "comma-delimited keywords to search by")
	filtersPtr := flag.String("filters", "", "comma-delimited (inclusive) filters for dietary information")
	xfiltersPtr := flag.String("xfilters", "", "comma-delimited (exclusive) filters for dietary information")
	flag.Parse()
	if *serverModePtr {
		fmt.Println("Server mode")
		server := NewResponseServer()

		fmt.Println("Serving on http://localhost:" + portNum)
		if err := http.ListenAndServe(":" + portNum, server); err != nil {
			log.Fatalf("could not listen on port 4242 %v", err)
		}
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
		dayData := fetchDayData(*datePtr, keywords, filters, xfilters)
		serializeMealData(dayData.Breakfast)
		serializeMealData(dayData.Lunch)
		serializeMealData(dayData.Dinner)
	}
}