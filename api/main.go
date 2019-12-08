package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const PORT = ":8000"

func main() {
	http.HandleFunc("/api/info", Info)
	hostname := os.Getenv("HOSTNAME")
	log.Printf("%s listening on %s\n", hostname, PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}

func Info(w http.ResponseWriter, r *http.Request) {
	type Response struct{ Status string }
	log.Printf("http %s\n", r.RequestURI)
	sendJSON(w, Response{Status: "OK"})
}

func sendJSON(w http.ResponseWriter, o interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(o); err != nil {
		log.Print(err)
	}
}
