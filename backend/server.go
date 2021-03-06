package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Post struct {
	User    string
	Thereds []string
}

type Post2 struct {
	Id      int
	Content string
	Author  string
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Print("test API")
	fmt.Fprintf(writer, "HelloWorld,%s!", request.URL.Path[1:])
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
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

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "go test",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "go test2",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=postgres host=localhost port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
}

//参照
func Posts() (posts []Post2, err error) {
	rows, err := db.Query("select id, content, author from posts")
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post2{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := Posts()
	if err != nil {
		return
	}
	var post Post2
	json.Unmarshal(body, &post)
	json := make([]byte, 1)
	w.Write([]byte(body))
}


//新規投稿
func (post *Post2) Create() (err error) {
	statement := "insert into posts(content,author) values($1,$2) returning id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post2
	json.Unmarshal(body, &post)
	err := post.Create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/v1/hello", handler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/redirect", redirectHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/cookie", setCookieHandler)
	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/post", postHandler)
	

	//db
	// post := Post2{Content: "hello", Author: "test man"}
	// fmt.Println(post)
	// post.Create()
	// fmt.Println(post)

	fmt.Println(Posts())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
