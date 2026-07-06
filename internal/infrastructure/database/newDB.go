package database

import (
	"projeto-golang/internal/domain/campaign"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "host=localhost user=emailn_dev password=luis1407 dbname=email_campaign port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Something happens wrong")
	}
	// Para criar referência de chave estrangeira
	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db
}
