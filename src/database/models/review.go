package models

type Review struct {
	ID     int64
	BookID int64
	Rating int8
	Text   string
	UserID int64
}
