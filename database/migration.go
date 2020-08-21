package database

import (
	"url-shortener/links"
	"url-shortener/users"

	"github.com/jinzhu/gorm"
)

// Migrate is the method for database migrations
func Migrate(db *gorm.DB) (err error) {
	db.AutoMigrate(links.Links{}, users.Users{})

	db.Model(&links.Links{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	return
}
