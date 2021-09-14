package migrations

import (
	"estudar.com.br/entity"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(entity.Book{})
}
