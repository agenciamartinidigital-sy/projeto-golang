package database

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func newDB() *gorm.DB {
	dsn := "user:luis1407@tcp(127.0.0.1:3306)/email_campaign?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect to database")
	}
	return db
}
