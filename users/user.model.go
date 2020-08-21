package users

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Users is the model for user table
type Users struct {
	gorm.Model
	Name     string `gorm:"column:name; type:varchar(100); not null;" json:"name"`
	Email    string `gorm:"column:email; type:varchar(255); not null;" json:"email"`
	Password string `gorm:"column:password; type:varchar(255); not null;" json:"password"`
}

func (u *Users) BeforeSave(scope *gorm.Scope) (err error) {
	if password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost); err == nil {
		scope.SetColumn("Password", password)
	}

	return
}
