package main

import (
	"fmt"
	"log"
	"net/http"
)

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

func main() {
	http.HandleFunc("/v1/hello", handler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/redirect", redirectHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
