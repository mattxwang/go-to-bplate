package main

import (
	"encoding/json"
    // "fmt"
    "net/http"
)

const jsonContentType = "application/json"

type ResponseServer struct {
	http.Handler
}

func NewResponseServer() *ResponseServer {
	s := new(ResponseServer)
	router := http.NewServeMux()
	router.Handle("/today", http.HandlerFunc(s.todayHandler))
	s.Handler = router
	return s
}

func (p *ResponseServer) todayHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	dayData := fetchDayData("Today", []string{}, []string{}, []string{})
    json.NewEncoder(w).Encode(dayData)
}