package users

import (
	"github.com/jinzhu/gorm"
)

type usersRepo struct {
	DB *gorm.DB
}

type UsersRepository interface {
	User(id int) Users
	Users() []Users
	Create(u Users) Users
	Update(id int, u Users) Users
	Delete(id int) bool
	FindUserByEmail(email string) Users
}

func NewUsersRepo(DB *gorm.DB) UsersRepository {
	return &usersRepo{DB}
}

func (r *usersRepo) User(id int) Users {
	var u Users

	r.DB.Select([]string{"ID", "Name", "Email"}).Find(&u, id)

	return u
}

func (r *usersRepo) Users() []Users {
	var uu []Users

	r.DB.Select([]string{"ID", "Name", "Email"}).Find(&uu)

	return uu
}

func (r *usersRepo) Create(u Users) Users {
	r.DB.Create(&u)

	return u
}

func (r *usersRepo) Update(id int, u Users) Users {
	var user Users

	r.DB.Model(&user).Updates(&u)

	return user
}

func (r *usersRepo) Delete(id int) bool {
	var u Users

	err := r.DB.Delete(&u, id).Error

	return err == nil
}

func (r *usersRepo) FindUserByEmail(email string) Users {
	var u Users

	r.DB.Where("LOWER(email) = ?", email).Find(&u)

	return u
}
