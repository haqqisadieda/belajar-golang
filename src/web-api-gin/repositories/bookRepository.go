package repositories

import (
	e "golang/src/web-api-gin/entities"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]e.Book, error)
	FindById(id int) (e.Book, error)
	Create(book e.Book) (e.Book, error)
	Update(book e.Book) (e.Book, error)
	Delete(book e.Book) (e.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]e.Book, error) {
	var books []e.Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindById(id int) (e.Book, error) {
	var book e.Book

	err := r.db.Find(&book, id).Error

	return book, err
}

func (r *repository) Create(book e.Book) (e.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) Update(book e.Book) (e.Book, error) {
	err := r.db.Save(&book).Error

	return book, err
}

func (r *repository) Delete(book e.Book) (e.Book, error) {
	err := r.db.Delete(&book).Error

	return book, err
}
