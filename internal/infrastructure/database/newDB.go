package database

import (
	"os"
	"projeto-golang/internal/domain/campaign"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Something happens wrong")
	}
	// Para criar referência de chave estrangeira
	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db
}
