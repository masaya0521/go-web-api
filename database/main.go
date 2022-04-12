package main // mainパッケージであることを宣言

import (
	// fmtモジュールをインポート

	"database/sql"
	"fmt"
	"log"
)

type Users struct {
	id     string `db:"id"`
	t_name string `db:"t_name"`
}

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal("OpenError: ", err)
	}
	defer db.Close()

	//データの確認
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var u Users
	for rows.Next() {
		err := rows.Scan(&u.id, &u.t_name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %s, Name: %s\n", u.id, u.t_name)
	}

	// データの挿入
	/*
	   id := '2'
	   t_name := "2taro"
	   _,err = db.Exec("INSERT INTO users(id, t_name) VALUES($1,$2)", id, t_name )
	   if err != nil {
	       fmt.Println(err)
	   }
	*/

	if err := db.Ping(); err != nil {
		log.Fatal("PingError: ", err)
	}
}
