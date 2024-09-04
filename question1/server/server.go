package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Define the handler function for the endpoint
func getData(w http.ResponseWriter, r *http.Request) {
	// Read the file data
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error reading file:", err)
		return
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Write the file data to the response
	w.Write(data)
}

func main() {
	http.HandleFunc("/producers", getData)
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
