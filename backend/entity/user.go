package entity

type User struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
}