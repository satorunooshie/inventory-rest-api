package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type Item struct {
	Name string `json:"Title"`
	Desc string `json:"Desc"`
	Price float64 `json:"Price"`
}

type Inventory []Item

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Endpoint")
}

func getInventory(w http.ResponseWriter, r *http.Request) {
	inventory := Inventory{
		Item{Name: "Cheese", Desc: "A fine block of cheese", Price: 400},
	}
	fmt.Println("Endpoint Called: getInventory()")
	json.NewEncoder(w).Encode(inventory)
}

func handleRequest() {
router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	log.Fatal(http.ListenAndServe(":1000", router))
	/*
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":1000", nil))
	 */
}

func main() {
	handleRequest()
}