package entities

type Book struct {
	ID           int64
	Title        string
	Description  string
	Author       string
	Cover        string
	Price        int32
	CategoryName string
	CategoryID   int64
	Reviews      []Review
}
