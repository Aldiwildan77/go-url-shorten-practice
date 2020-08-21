package links

import (
	"math/rand"

	"github.com/jinzhu/gorm"
)

type linksRepo struct {
	DB *gorm.DB
}

type LinksRepository interface {
	Link(id int) Links
	Links() []Links
	Create(l Links) Links
	Update(id int, l Links) Links
	Delete(id int) bool
	FindByShortenLink(link string) Links
	GenerateShortenLink() string
}

func NewLinksRepo(DB *gorm.DB) LinksRepository {
	return &linksRepo{DB}
}

func (r *linksRepo) Link(id int) Links {
	var l Links

	r.DB.Find(&l, id)

	return l
}

func (r *linksRepo) Links() []Links {
	var ll []Links

	r.DB.Find(&ll)

	return ll
}

func (r *linksRepo) Create(l Links) Links {
	r.DB.Create(&l)

	return l
}

func (r *linksRepo) Update(id int, l Links) Links {
	var link Links

	r.DB.Model(&link).Update(&l)

	return l
}

func (r *linksRepo) Delete(id int) bool {
	var l Links

	err := r.DB.Delete(&l, id).Error

	return err == nil
}

func (r *linksRepo) FindByShortenLink(link string) Links {
	var l Links

	r.DB.Find(&l).Where("translated_url = ?", link)

	return l
}

func (r *linksRepo) GenerateShortenLink() string {
	var letterRunes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var count int8 = 5

	chunkOfString := make([]byte, count)

	for i := 0; i < int(count); i++ {
		chunkOfString[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(chunkOfString)
}
