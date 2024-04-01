package main

import (
	"crypto/md5"
	"encoding/hex"
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

// shortenURL shortens a long incoming URL and returns the short URL.
func shortenURL(w http.ResponseWriter, r *http.Request) {

	// Getting the original URL from the form data.
	originalURL := r.FormValue("url")

	//Creating a hash to shorten the URL.
	hash := hashURL(originalURL)

	// A short URL is created by taking the first 6 characters of the hash.
	shortURL := hash[:6]

	// Matching the generated short URL with the original URL.
	urlMap[shortURL] = originalURL

	// Returning the short URL to the user.
	fmt.Fprintf(w, "Shortened URL: http://localhost:8080/%s", shortURL)
}

// redirectURL redirects a short URL to the original URL.
func redirectURL(w http.ResponseWriter, r *http.Request) {

	// Getting the short URL from the URL.
	shortURL := mux.Vars(r)["shortURL"]

	// Getting the original URL matched to the short URL.
	originalURL, ok := urlMap[shortURL]
	if !ok {
		// If no match is found, a 404 Error is returned.
		http.NotFound(w, r)
		return
	}

	// The user is redirected from the short URL to the original URL.
	http.Redirect(w, r, originalURL, http.StatusFound)
}

// hashURL converts a URL to a hash using a hashing algorithm.
func hashURL(url string) string {

	// Creating a hasher to use the MD5 hashing algorithm.
	hasher := md5.New()

	// Writing the bytes of the URL to the hasher.
	hasher.Write([]byte(url))

	// Calculating and converting the hash.
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
