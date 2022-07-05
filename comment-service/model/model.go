package model

type Comment struct {
	Id      string `json:"id"`
	User    string `json:"user"`
	Content string `json:"content"`
}
