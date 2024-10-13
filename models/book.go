package models

type Book struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishYear uint   `json:"publishYear"`
	Gender      string `json:"gender"`
}
