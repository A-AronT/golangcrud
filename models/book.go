package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Bookcases []Bookcase `gorm:"many2many:book_bookcases;"`
}
