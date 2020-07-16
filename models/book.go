package models

type Book struct {
	ID            uint        `json:"id" gorm:"primary_key"`
	Title         string      `json:"title"`
	Author        string      `json:"author"`
	PublicationID Publication `json:"publication_id" gorm:"foreignkey:ID"`
}
