package main

import (
	"database/sql"
	"fmt"
	"time"

	"context"
	"log"

	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
)

func testEnvironmentVariable() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	test := os.Getenv("TEST")
	fmt.Printf("Test = [%s]", test)
}

func getLanguages() []string {
	return []string{"C#", "Java", "Ruby", "Python", "JavaScript", "Go", "Rust", "TypeScript"}
}

type Skill struct {
	IDCode string
	Name  string
}

func getSkillsFromMongo() []Skill {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	mongo_conn := os.Getenv("MONGO_CONNECTION")
	mongo_database := os.Getenv("MONGO_DATABASE")
	mongo_collection := os.Getenv("MONGO_COLLECTION")

	clientOptions := options.Client().ApplyURI(mongo_conn)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database(mongo_database).Collection(mongo_collection)

	filter := bson.D{}

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

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

	var skills []Skill
	for _, result := range results {
		var idCode, name string
		if result["idCode"] != nil {
			idCode = result["idCode"].(string)
		}
		if result["name"] != nil {
			name = result["name"].(string)
		}
		skill := Skill{
			IDCode: idCode,
			Name:  name,
		}
		skills = append(skills, skill)
	}

	return skills

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
