package main

import (
	"encoding/json"
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

// Global var of Animal Slice (This is only for this small app, as we're not using an external DB)
// TODO: Implement DB
var animals []Animal

// Functions
// Fetch All Animals
func getAnimals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animals)
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
	// *Sample Data (This is only due to the fact we're not using a DB for this project)
	// TODO: Implement DB
	animals = append(animals, Animal{ID: "1", Name: "Hector", Diet: "Fruits", Species: &Species{Name: "Chimp", Family: "Mammal", Domestication: "Domestic", Endangered: "No"}})
	animals = append(animals, Animal{ID: "2", Name: "Lauritz", Diet: "Fruits", Species: &Species{Name: "Giraffe", Family: "Mammal", Domestication: "Wild", Endangered: "Yes"}})
	animals = append(animals, Animal{ID: "3", Name: "Karl", Diet: "Meat", Species: &Species{Name: "Dog", Family: "Canine", Domestication: "Domestic", Endangered: "No"}})

	// Create route endpoints + handlers - Create the CRUD handlers here (Using animals for this example, change to your usage)
	router.HandleFunc("/api/animals", getAnimals).Methods("GET")
	router.HandleFunc("/api/animals/{id}", getAnimal).Methods("GET")
	router.HandleFunc("/api/animals", createAnimal).Methods("POST")
	router.HandleFunc("/api/animals/{id}", updateAnimal).Methods("PUT")
	router.HandleFunc("/api/animals/{id}", deleteAnimal).Methods("DELETE")

	// Run the Server
	log.Fatal(http.ListenAndServe(":8000", router))
}
