package links

import "github.com/jinzhu/gorm"

// Links is the model for url table
type Links struct {
	gorm.Model
	OriginalURL   string `gorm:"column:original_url; type:varchar(100)" json:"original_url"`
	TranslatedURL string `gorm:"column:translated_url; type:varchar(255); unique_index"`
	UserID        uint   `gorm:"column:user_id; index; not null;" json:"user_id"`
}
