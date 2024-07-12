package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	port := 7788

	http.HandleFunc("/languages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]string{"C#", "Java", "Ruby", "Python", "JavaScript", "Go"})
	})
	fmt.Printf("listening at http://localhost:%v/languages\n", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}