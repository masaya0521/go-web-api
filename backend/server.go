package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "HelloWorld,%s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/v1/hello", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
