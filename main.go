package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type response struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func main() {
	http.HandleFunc("/time/RFC3339", timeHandler)
	fmt.Println("Server ishga tushdi http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	if r.Header.Get("Accept") == "application/json" {
		jsonResponse(w, currentTime)
		return
	}
	textResponse(w, currentTime)
}

func jsonResponse(w http.ResponseWriter, t time.Time) {
	rsp := response{
		DayOfWeek:  t.Weekday().String(),
		DayOfMonth: t.Day(),
		Month:      t.Month().String(),
		Year:       t.Year(),
		Hour:       t.Hour(),
		Minute:     t.Minute(),
		Second:     t.Second(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rsp)
}

func textResponse(w http.ResponseWriter, t time.Time) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, t.Format(time.RFC3339))
}
