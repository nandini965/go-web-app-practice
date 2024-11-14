package main

import (
	"log"
	"net/http"
)

// Existing route handlers
func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

// New catch-all handler for undefined routes
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func main() {
	// Define routes
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/courses", coursePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)
	// Catch-all route
	http.HandleFunc("/", notFoundHandler)

	// Start the server
	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal(err)
	}
}
