package db

import (
	"lld/blogging/model"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=blogging port=5000 sslmode=disable TimeZone=Asia/Shanghai"
	var err error

	once.Do(func() {
		// db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	})
	if err != nil {
		log.Fatal("failed to connec with DB", err)
	}
	db.AutoMigrate(&model.User{}, &model.Blog{}, &model.Comment{}, &model.Subscription{})
	return db
}
