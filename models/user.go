package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Phone     string
	Email     string `gorm:"unique"`
	Password  string `gorm:"unique"`
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
