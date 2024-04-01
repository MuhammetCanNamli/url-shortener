package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// urlMap, a map that maps short URLs to original URLs.
var urlMap map[string]string

func main() {
	//creating the urlMap.
	urlMap = make(map[string]string)

	// Creating a new router.
	r := mux.NewRouter()

	// Assigning the router for the main root directory.
	http.Handle("/", r)

	// The server starts listening on port 8080.
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func shortenURL(w http.ResponseWriter, r *http.Request) {

}

func redirectURL(w http.ResponseWriter, r *http.Request) {

}

func hashURL(url string) {

}
