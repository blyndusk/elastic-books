package models

type Book struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Resume string `json:"resume"`
}

type Books []Book
