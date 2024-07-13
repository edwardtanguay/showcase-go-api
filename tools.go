package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func getLanguages() []string {
	return []string{"C#", "Java", "Ruby", "Python", "JavaScript", "Go", "Rust", "TypeScript"}
}

func getHowtos() {

	// Open the database
	db, err := sql.Open("sqlite3", "./data/main.sqlite")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Query the database
	rows, err := db.Query("SELECT id, category, title FROM howtos")
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var id int
		var category string
		var title string
		err = rows.Scan(&id, &category, &title)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("id: %d, category: %s, title: %s\n", id, category, title)
	}

	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		fmt.Println("Error iterating over rows:", err)
		return
	}
}
