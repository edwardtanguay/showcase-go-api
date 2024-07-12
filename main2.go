package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	fmt.Println("Listening on http://localhost:3333")
	http.ListenAndServe(":3333", nil)
}
