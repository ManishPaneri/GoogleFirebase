package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/demo", demoHandler)
	handler = c.Handler(handler)
	http.ListenAndServe(":9000", handler)
}

// Demo Collection Call
func demoHandler(w http.ResponseWriter, r *http.Request) {

	data := controllers.MapCollectionUrl(w, r)
	//fmt.Println("response :", data)
	returnJson, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnJson)
}
