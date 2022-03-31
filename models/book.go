package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Authors []Author `gorm:"many2many:book_authors;"`
}
