package main

import (
	"fmt"
	"log"
	"net/http"
)

func startHandlers() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./html/home.html")
	})

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/comment", commentHandler)
	http.HandleFunc("/like", likeHandler)
	http.HandleFunc("/filter", filterHandler)
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/logout", logoutHandler)
}

func main() {

	dbStart()
	defer myDb.Close()

	http.Handle("/html/", http.FileServer(http.Dir("./")))

	startHandlers()

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
