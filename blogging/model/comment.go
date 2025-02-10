package model

type Comment struct {
	Id       int    `gorm:"primaryKey"`
	Content  string `gorm:"not null"`
	UserID   uint
	BlogID   uint
	ParentID *uint
}
