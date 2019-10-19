package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const jsonContentType = "application/json"

type ResponseServer struct {
	cache map[string]*MenuHit
	http.Handler
}

type MenuHit struct {
	data *DayData
}

func NewResponseServer() *ResponseServer {
	s := new(ResponseServer)
	s.cache = map[string]*MenuHit{}
	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(s.todayHandler))
	router.Handle("/today", http.HandlerFunc(s.todayHandler))
	router.Handle("/tomorrow", http.HandlerFunc(s.tomorrowHandler))
	router.Handle("/date/", http.HandlerFunc(s.dateEndpointHandler))
	s.Handler = router
	return s
}

func NewMenuHit(data *DayData) *MenuHit {
	m := new(MenuHit)
	m.data = data
	return m
}

func getMenuHit(p *ResponseServer, dateString string) *MenuHit {
	log.Println("Getting all meals for " + dateString)
	hit, exists := p.cache[dateString]
	if !exists {
		log.Println("Cache miss, retrieving from server")
		newDayData := fetchDayData(dateString)
		log.Println("Server responded, caching result for future")
		newMenuHit := NewMenuHit(newDayData)
		p.cache[dateString] = newMenuHit
		hit = newMenuHit
	} else {
		log.Println("Cache hit!")
	}
	return hit
}

func clearCache(p *ResponseServer) {
	p.cache = map[string]*MenuHit{}
	log.Println("Cache reset.")
}

func fetchNextWeekMenu(p *ResponseServer) {
	currentTime := time.Now()
	for i := 0; i < 7; i++ {
		dateString := currentTime.AddDate(0, 0, i).Format("2006-01-02")
		getMenuHit(p, dateString)
	}
}

func updateCache(p *ResponseServer) {
	log.Println("Updating the cache!")
	clearCache(p)
	fetchNextWeekMenu(p)
}

func populateSearchOptions(r *http.Request) *SearchOptions {
	searchOptions := new(SearchOptions)

	keywordsQuery, exists := r.URL.Query()["keywords"]
	if exists && len(keywordsQuery[0]) > 0 {
		searchOptions.keywords = splitStringsByComma(keywordsQuery[0])
	}
	filtersQuery, exists := r.URL.Query()["filters"]
	if exists && len(filtersQuery[0]) > 0 {
		searchOptions.filters = splitStringsByComma(filtersQuery[0])
	}
	xfiltersQuery, exists := r.URL.Query()["xfilters"]
	if exists && len(xfiltersQuery[0]) > 0 {
		searchOptions.xfilters = splitStringsByComma(xfiltersQuery[0])
	}
	return searchOptions
}

func genericDateHandler(p *ResponseServer, w http.ResponseWriter, r *http.Request, dateString string) {
	searchOptions := populateSearchOptions(r)
	hit := getMenuHit(p, dateString)
	filteredData := filterDayData(hit.data, searchOptions)
	json.NewEncoder(w).Encode(filteredData)
}

func (p *ResponseServer) todayHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	currentTime := time.Now()
	dateString := currentTime.Format("2006-01-02")
	genericDateHandler(p, w, r, dateString)
}

func (p *ResponseServer) tomorrowHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	currentTime := time.Now()
	dateString := currentTime.AddDate(0, 0, 1).Format("2006-01-02")
	genericDateHandler(p, w, r, dateString)
}

func (p *ResponseServer) dateEndpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	dateString := r.URL.Path[len("/date/"):]
	genericDateHandler(p, w, r, dateString)
}
