package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Endpoint")
}

func handleRequest() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":1000", nil))
}

func main() {
	handleRequest()
}