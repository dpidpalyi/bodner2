package main

import (
	"log"
	"net/http"
	"time"
)

func MiddleWare(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}

func GetTime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	t := time.Now().Format(time.RFC3339)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(t))
	return
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", MiddleWare(GetTime))

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	log.Fatal(server.ListenAndServe())
}
