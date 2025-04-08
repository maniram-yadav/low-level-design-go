package models

import "time"

type Answer struct {
	ID         uint      `json:"id"`
	Body       string    `json:"body"`
	UserID     uint      `json:"userId"`
	User       User      `json:"user,omitempty"`
	QuestionID uint      `json:"questionId"`
	Question   Question  `json:"-"`
	Comments   []Comment `json:"comments,omitempty"`
	IsAccepted User      `json:"isAccepted"`
	VoteCount  uint      `json:"voteCount"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
