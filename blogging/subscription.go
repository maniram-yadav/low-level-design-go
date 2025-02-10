package blogging

type Subscription struct {
	ID    int    `gorm:"primaryKey"`
	Email string `gorm:"unique;not null"`
}
