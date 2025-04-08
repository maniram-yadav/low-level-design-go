package models

import "time"

type Question struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserID    uint      `json:"userId"`
	User      User      `json:"user,omitempty"`
	Answers   []Answer  `json:"answers,omitempty"`
	Tags      []Tag     `json:"tags,omitempty" gorm:"many2many:question_tags;"`
	VoteCount int       `json:"voteCount"`
	Views     int       `json:"views"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Status    string    `json:"status"`
}
