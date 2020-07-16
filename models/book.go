package models

type Book struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublicationID uint   `json:"publication_id"`
}
