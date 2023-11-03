package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Firstname string `json:"firstname" gorm:"not null"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email" gorm:"not null;unique"`
	Birthday  string `json:"birthday" gorm:"not null"`
}
