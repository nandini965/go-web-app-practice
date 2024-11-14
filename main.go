package main

import (
	"log"
	"net/http"
)

// Home page handler
func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

// Courses page handler
func coursePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/courses.html")
}

// About page handler
func aboutPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

// Contact page handler
func contactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

// Root path handler
func rootPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusFound)
}

func main() {
	// Register the handlers
	http.HandleFunc("/", rootPage)          // Handle the root path
	http.HandleFunc("/home", homePage)      // Handle the /home path
	http.HandleFunc("/courses", coursePage) // Handle the /courses path
	http.HandleFunc("/about", aboutPage)    // Handle the /about path
	http.HandleFunc("/contact", contactPage) // Handle the /contact path

	// Start the server
	log.Println("Server starting on port 8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
