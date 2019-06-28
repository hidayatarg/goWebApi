package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Dec string `json:"desc"`
	Content string `json:"content"`

}
type Articles []Article

func allArticles (w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title:"Test Title", Dec:"Test Description", Content: "Hello World"},
	}
	fmt.Println("EndPoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homPage (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hompage Endpoint Hit")
}

func handleRequests () {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homPage)
	// only call with get method
	router.HandleFunc("/articles", allArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}


