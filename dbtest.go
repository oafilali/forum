package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var myDb *sql.DB

func dbStart() {
	var err error
	myDb, err = sql.Open("sqlite3", "./forumtest.db")
	errorCheck("Starting the database failed", err)
	
	err = db.Ping()
	errorCheck("Connecting to the database failed", err)

	log.Panicln("Database connected succesfully")
	startTabeles()
}

func startTabeles() {

	tables := map[string]string{
		"Users": `CREATE TABLE IF NOT EXISTS Users (
			user_id  INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			email  TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);`,
		"Posts": `CREATE TABLE IF NOT EXISTS POSTS (
			post_id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			CONTENT TEXT NOT NULL,
			FOREIGN KEY(user_id) REFERENCES Users(user_id)
		);`,
		"Comments": `CREATE TABLE IF NOT EXISTS Comments (
			comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT NOT NULL
			FOREIGN KEY(user_id) REFERENCES Users(user_id),
			FOREIGN KEY(post_id) REFERENCES Posts(post_id)
		);`,
		"Reactions": `CREATE TABLE IF NOT EXISTS Reactions (
			reaction_id INTEGER PRIMARY KEY AUTOINCREMENT,
			type TEXT CHECK(type IN ('like', 'deslike')),
			FOREIGN KEY(user_id) REFERENCES Users(user_id),
			FOREIGN KEY(post_id) REFERENCES Posts(post_id)
		);`,
		"Sesseions": `CREATE TABLE IF NOT EXISTS Sessions (
			
		);`

}

func errorCheck(msg string, err error) {
	if err != nil {
		log.Fatal(msg,err)
	}
}
