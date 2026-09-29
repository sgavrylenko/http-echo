package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func logRequest(r *http.Request, duration time.Duration) {
	log.Printf(
		"[%s] %s %s %s - Оброблено за %v",
		r.RemoteAddr,
		r.Method,
		r.URL.Path,
		r.Proto,
		duration,
	)
}

func echoHeadersHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(r.Header); err != nil {
		log.Printf("Помилка кодування JSON: %v", err)
	}

	logRequest(r, time.Since(startTime))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/", echoHeadersHandler)

	addr := fmt.Sprintf(":%s", port)

	log.Printf("Сервер запущено на порту %s...", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Помилка запуску сервера: %v", err)
	}
}
