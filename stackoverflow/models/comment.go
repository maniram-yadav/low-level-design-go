package models

import "time"

type Comment struct {
	ID         uint      `json:"id"`
	Text       string    `json:"text"`
	UserID     uint      `json:"userId"`
	User       User      `json:"user,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
	ParentID   uint      `json:"parentId"`
	ParentType string    `json:"parentType"`
}
