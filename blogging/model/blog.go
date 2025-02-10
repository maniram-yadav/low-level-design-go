package model

type Blog struct {
	Id      int    `gorm:"primarykey"`
	Title   string `gorm:"non null"`
	Content string `gorm:"not null"`
	UserID  uint
}
