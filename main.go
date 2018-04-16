package main

import (
	"GoogleFirebase/controllers"
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/demo", demoHandler)
	http.ListenAndServe(":9000", nil)
}

// Demo Collection Call
func demoHandler(w http.ResponseWriter, r *http.Request) {

	data := controllers.MapCollectionUrl(w, r)
	returnJson, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnJson)
}
