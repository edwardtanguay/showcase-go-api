package main

import (
	"database/sql"
	"fmt"
	"time"

	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getLanguages() []string {
	return []string{"C#", "Java", "Ruby", "Python", "JavaScript", "Go", "Rust", "TypeScript"}
}

func getTodosWithMongo() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Specify the database and collection
	collection := client.Database("book-app-234").Collection("books")

	// Define a filter to match documents
	filter := bson.D{}

	// Find multiple documents
	var results []bson.M
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	// Print the results
	for _, result := range results {
		fmt.Printf("Found document: %v\n", result)
	}

	// Close the connection once done
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}

func getHowtosWithSqlite() {

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
