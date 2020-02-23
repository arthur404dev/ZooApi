package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Animal Structs
// Animal Model
type Animal struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Species *Species `json:"species"`
	Diet    string   `json:"diet"`
}

// Species Model
type Species struct {
	Name          string `json:"name"`
	Family        string `json:"family"`
	Domestication string `json:"domestication"`
	Endangered    string `json:"endangered"`
}

// Functions
// Fetch All Animals
func getAnimals(w http.ResponseWriter, r *http.Request) {

}

// Fetch Single Animal
func getAnimal(w http.ResponseWriter, r *http.Request) {

}

// Create a New Animal
func createAnimal(w http.ResponseWriter, r *http.Request) {

}

// Update a Single Animal
func updateAnimal(w http.ResponseWriter, r *http.Request) {

}

// Delete Animal
func deleteAnimal(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Start the router for the main endpoint - Going to use Mux for this project
	router := mux.NewRouter()

	// Create route endpoints + handlers - Create the CRUD handlers here (Using animals for this example, change to your usage)
	router.HandleFunc("/api/animals", getAnimals).Methods("GET")
	router.HandleFunc("/api/animals/{id}", getAnimal).Methods("GET")
	router.HandleFunc("/api/animals", createAnimal).Methods("POST")
	router.HandleFunc("/api/animals/{id}", updateAnimal).Methods("PUT")
	router.HandleFunc("/api/animals/{id}", deleteAnimal).Methods("DELETE")

	// Run the Server
	log.Fatal(http.ListenAndServe(":6000", router))
}
