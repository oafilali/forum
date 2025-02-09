package main

import (
	"fmt"
	"log"
	"net/http"
)

func main2() {
	// Initialize the database
	initDB()
	defer db.Close()

	// Serve static files (CSS)
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))

	// Serve the home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./html/home.html")
	})

	// Register HTTP handlers
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/comment", commentHandler)
	http.HandleFunc("/like", likeHandler)
	http.HandleFunc("/filter", filterHandler)
	http.HandleFunc("/posts", postsHandler) // New route to display posts
	http.HandleFunc("/logout", logoutHandler) // New route to handle logout

	// Start the server
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}