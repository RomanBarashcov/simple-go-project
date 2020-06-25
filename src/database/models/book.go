package models

type Book struct {
	ID          int64
	Title       string
	Description string
	Author      string
	Cover       string
	Price       int32
	CategoryID  int64
}
