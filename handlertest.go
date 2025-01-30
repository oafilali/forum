package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func errorCheckHandlers(w http.ResponseWriter, msg string, err error, code int) bool {
	if err != nil {
		http.Error(w, msg, code)
		log.Println(msg, err)
		return true
	}
	return false
}

func checkEmailCount(w http.ResponseWriter, emailCount int) bool {
	if emailCount > 1 {
		http.Error(w, "Database integrity error: Duplicate emails found", http.StatusInternalServerError)
		return true
	}
	if emailCount == 1 {
		http.Error(w, "Email already taken", http.StatusBadRequest)
		return true
	}
	return false
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		var emailCount int
		err := myDb.QueryRow("SELECT COUNT(*) FROM USERS WHERE EMAIL = ?", email).Scan(&emailCount)
		if errorCheckHandlers(w, "Database error", err, http.StatusInternalServerError) {
			return
		}
		if checkEmailCount(w, emailCount) {
			return
		}

		_, err = myDb.Exec("INSERT INTO Users (username, email, password) VALUES (?, ?, ?)", username, email, password)
		if errorCheckHandlers(w, "Failed to register user", err, http.StatusInternalServerError) {
			return
		}

		fmt.Fprintf(w, "User registered successfully")
		log.Println("User registered successfully")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		http.ServeFile(w, r, "./html/register.html")
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		var myDbPassword string
		var user_id int
		err := myDb.QueryRow("SELECT user_id, password FROM Users WHERE email = ?", email).Scan(&user_id, &myDbPassword)
		if errorCheckHandlers(w, "User not found", err, http.StatusNotFound) {
			return
		}
		if password != myDbPassword {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		sessionID, err := uuid.NewV4()
		if errorCheckHandlers(w, "Failed to generate session ID", err, http.StatusInternalServerError) {
			return
		}

		created_at := time.Now()
		expires_at := time.Now().Add(2 * time.Hour)

		_, err = myDb.Exec("INSERT INTO Sessions (session_id, user_id, created_at, expires_at) VALUES (?, ?, ?, ?)", sessionID.String(), user_id, created_at, expires_at)
		if errorCheckHandlers(w, "Failed to create session", err, http.StatusInternalServerError) {
			return
		}

		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    sessionID.String(),
			Expires:  expires_at,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)

		fmt.Fprintf(w, "Login successful")
		log.Println("Login successful")
		http.Redirect(w, r, "/posts", http.StatusSeeOther)
	} else {
		http.ServeFile(w, r, "./html/login.html")
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")
		category := r.FormValue("category")

		_, err := myDb.Exec("INSERT INTO Posts ")
	}

}
