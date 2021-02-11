package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type Item struct {
	UID string `json:"UID"`
	Name string `json:"Name"`
	Desc string `json:"Desc"`
	Price float64 `json:"Price"`
}

var inventory []Item

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Function Called: home()")
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Function Called: getInventory()")
	json.NewEncoder(w).Encode(inventory)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	// obtain item from request JSON
	_ = json.NewDecoder(r.Body).Decode(&item)
	inventory = append(inventory, item)
	// Show item in response JSON for
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params)
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory/{uid}", updateItem).Methods("GET")
	router.HandleFunc("/inventory", addItem).Methods("POST")
	log.Fatal(http.ListenAndServe(":1000", router))
	/*
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":1000", nil))
	 */
}

func main() {
	// Data Store
	inventory = append(inventory, Item{
		UID: "0",
		Name: "Cheese",
		Desc: "A fine block of cheese",
		Price: 400,
	})
	handleRequest()
}