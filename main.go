package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get the Params
	// Loop through all entries and find the ocurrence
	for _, item := range animals {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Animal{})
}

// Create a New Animal
func createAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var animal Animal
	_ = json.NewDecoder(r.Body).Decode(&animal)
	animal.ID = strconv.Itoa(rand.Intn(10000000)) // This is an not Hashed ID, only for testing purpose - This is not Safe for production!
	animals = append(animals, animal)
	json.NewEncoder(w).Encode(animal)
}

// Update a Single Animal
func updateAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range animals {
		if item.ID == params["id"] {
			animals = append(animals[:index], animals[index+1:]...)
			var animal Animal
			_ = json.NewDecoder(r.Body).Decode(&animal)
			animal.ID = params["id"]
			animals = append(animals, animal)
			json.NewEncoder(w).Encode(animal)
			return
		}
	}
	json.NewEncoder(w).Encode(animals)
}

// Delete Animal
func deleteAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range animals {
		if item.ID == params["id"] {
			animals = append(animals[:index], animals[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(animals)
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
