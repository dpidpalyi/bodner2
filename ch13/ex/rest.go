package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type JSONTime struct {
	Day_of_week  string `json:"day_of_week"`
	Day_of_month int    `json:"day_of_month"`
	Month        string `json:"month"`
	Year         int    `json:"year"`
	Hour         int    `json:"hour"`
	Minute       int    `json:"minute"`
	Second       int    `json:"second"`
}

func MiddleWare(sl *slog.Logger, next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sl.Debug("debug message", "ip", r.RemoteAddr)
		next(w, r)
	}
}

func GetTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	t := time.Now()
	w.WriteHeader(http.StatusOK)
	var out []byte
	if r.Header.Get("Accept") == "application/json" {
		out = buildJSON(t)
	} else {
		out = buildText(t)
	}
	w.Write(out)
	return
}

func buildJSON(now time.Time) []byte {
	t := JSONTime{
		Day_of_week:  now.Weekday().String(),
		Day_of_month: now.Day(),
		Month:        now.Month().String(),
		Year:         now.Year(),
		Hour:         now.Hour(),
		Minute:       now.Minute(),
		Second:       now.Second(),
	}
	out, _ := json.Marshal(t)
	return out
}

func buildText(now time.Time) []byte {
	return []byte(now.Format(time.RFC3339))
}

func main() {
	mux := http.NewServeMux()
	options := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(os.Stderr, options)
	mySlog := slog.New(handler)
	mux.HandleFunc("/", MiddleWare(mySlog, GetTime))

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	log.Fatal(server.ListenAndServe())
}
