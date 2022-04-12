package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	User    string
	Thereds []string
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "HelloWorld,%s!", request.URL.Path[1:])
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintf(w, "HelloWorld!")
}

func redirectHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Location", "https://www.goole.com")
	w.WriteHeader(302)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "user",
		Thereds: []string{"1", "2", "3"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	http.HandleFunc("/v1/hello", handler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/redirect", redirectHandler)
	http.HandleFunc("/json", jsonHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
