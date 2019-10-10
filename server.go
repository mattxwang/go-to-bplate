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
	router.Handle("/today", http.HandlerFunc(s.todayHandler))
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
		newDayData := fetchDayData(dateString, []string{}, []string{}, []string{})
		log.Println("Server responded, caching result for future")
		newMenuHit := NewMenuHit(newDayData)
		p.cache[dateString] = newMenuHit
		hit = newMenuHit
	} else {
		log.Println("Cache hit!")
	}
	return hit
}

func (p *ResponseServer) todayHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	dateString := time.Now().Format("2006-01-02")
	hit := getMenuHit(p, dateString)
	json.NewEncoder(w).Encode(hit.data)
}
