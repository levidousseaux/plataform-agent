package entity

import "time"

type User struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
