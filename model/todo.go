package model

import "github.com/jinzhu/gorm"

// Todo is a representative of `todos`
type Todo struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Completed int
}

// TodoSchema is a structure that uses by response payload
type TodoSchema struct {
	ID        uint
	Title     string
	Completed bool
}
