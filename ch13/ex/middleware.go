package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

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
	t := time.Now().Format(time.RFC3339)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(t))
	return
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
