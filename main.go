package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func Hello(name string) string {
	return "Hello, " + name
}

func main() {

	viper.SetConfigFile("ENV")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	port := fmt.Sprint(viper.Get("PORT"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	// http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
	// 	book := r.URL.Query().Get("book")
	// 	fmt.Printf("book: %q", book)
	// 	if book == "" {
	// 		fmt.Fprintf(w, "You must specify a book")
	// 	}

	// 	fmt.Fprintf(w, "Trying to get %s!", book)
	// })

	type Book struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var book Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		// print locally
		fmt.Printf("Title: %q, Author: %q\n", book.Title, book.Author)
		// send response
		fmt.Fprintf(w, "Successfully received book: %q by %q", book.Title, book.Author)
	})

	// recieve a name in json format and return a greeting in json format
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Printf("name: %q", name)
		if name == "" {
			name = "World"
		}
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
