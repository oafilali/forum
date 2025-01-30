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
	errorCheck("Starting the database failed: ", err)

	err = myDb.Ping()
	errorCheck("Connecting to the database failed: ", err)

	log.Println("Database connected succesfully")
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
			user_id INTEGER NOT NULL,
			title TEXT NOT NULL,
			CONTENT TEXT NOT NULL,
			FOREIGN KEY(user_id) REFERENCES Users(user_id)
		);`,
		"Comments": `CREATE TABLE IF NOT EXISTS Comments (
			comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			post_id INTEGER NOT NULL,
			content TEXT NOT NULL,
			FOREIGN KEY(user_id) REFERENCES Users(user_id),
			FOREIGN KEY(post_id) REFERENCES Posts(post_id)
		);`,
		"Reactions": `CREATE TABLE IF NOT EXISTS Reactions (
			reaction_id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
		    post_id INTEGER NOT NULL,
			type TEXT CHECK(type IN ('like', 'dislike')),
			FOREIGN KEY(user_id) REFERENCES Users(user_id),
			FOREIGN KEY(post_id) REFERENCES Posts(post_id)
		);`,
		"Sessions": `CREATE TABLE IF NOT EXISTS Sessions (
			session_id TEXT PRIMARY KEY,
			user_id INTEGER NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			expires_at DATETIME,
			FOREIGN KEY(user_id) REFERENCES Users(user_id)
		);`,
		"Categories": `CREATE TABLE IF NOT EXISTS Categories (
			category_id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE
		);`,
		"Post_Categories": `CREATE TABLE IF NOT EXISTS Post_Categories (
			post_id INTEGER NOT NULL,
    		category_id INTEGER NOT NULL,
			PRIMARY KEY(post_id, category_id),
			FOREIGN KEY(post_id) REFERENCES Posts(post_id),
			FOREIGN KEY(category_id) REFERENCES Categories(category_id)
		);`,
	}

	for _, table := range tables {
		_, err := myDb.Exec(table)
		errorCheck("Failed to create table: ", err)
	}

}

func errorCheck(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
