package models

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Resume string `json:"resume"`
}

type Books []Book
