package entity

import (
	"time"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	CreatedAt time.Time
}

// type Book struct {
// 	Id        uint
// 	Title     string `gorm:"size:128"`
// 	Category  int
// 	Author    string `gorm:"size:64"`
// 	CreatedAt time.Time
// }
