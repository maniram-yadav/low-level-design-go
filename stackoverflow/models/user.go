package models

import "time"

type User struct {
	ID           uint      `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Reputation   int       `json:"reputation"`
	CreatedAt    time.Time `json:"createdAt"`
	Badges       []Badge   `json:"badges,omitempty" gorm:"many2many:user_badges;"`
}
