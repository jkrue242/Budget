package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
)

type Message struct{ Greeting string `json:"greeting"` }

var staticFS embed.FS

func apiGreeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Greeting: "Hello from Go backend!"})
}

func main() {
	http.HandleFunc("/api/greeting", apiGreeting)

	dist, _ := fs.Sub(staticFS, "webclient/dist")
	http.Handle("/", http.FileServer(http.FS(dist)))

	log.Println("✅  Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
