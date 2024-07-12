package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"os"
)

func main() {

	test := os.Getenv("TEST")
	fmt.Printf("Test is: [%v]\n", test)

	port := 7788

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Showcase Go API</title>
	<style>
		* {
			margin: 0;
			padding: 0;
			box-sizing: border-box;
		}

		body {
			background-color: #ccc;
			margin: 1rem;
			font-family: monospace;
			line-height: 2rem;
		}

		ul {
			margin-left: 2rem;
		}
	</style>
</head>
<body>
	<h1>Showcase Go API</h1>	
	<p>This API is written in Go.</p>
	<p>It is being served on Debian with Nginx and PM2.</p>
	<p>Current available routes:</p>
	<ul>
		<li><a href="/languages">/languages</a> - a JSON array of computer languages</li>
	</ul>
</body>
</html>
		`)
	})

	http.HandleFunc("/languages", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]string{"C#", "Java", "Ruby", "Python", "JavaScript", "Go"})
	})

	fmt.Printf("listening at http://localhost:%v\n", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
