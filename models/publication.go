package models

type Publication struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Title string `json:"title" gorm:"unique;not null"`
}
