package main

import (
	"fmt"
	"log"
	"net/http"
)

func homPage (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hompage Endpoint Hit")
}

func handleRequests () {
	http.HandleFunc("/", homPage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}


