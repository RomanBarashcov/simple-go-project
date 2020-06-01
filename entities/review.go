package entities

type Review struct {
	ID     int
	Text   string
	Rating int8
	User   User
}
