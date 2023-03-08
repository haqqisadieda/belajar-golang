package services

import (
	e "golang/src/web-api-gin/entities"
	m "golang/src/web-api-gin/models"
	r "golang/src/web-api-gin/repositories"
)

type Service interface {
	FindAll() ([]e.Book, error)
	FindById(id int) (e.Book, error)
	Create(bookRequest m.BookCreateRequest) (e.Book, error)
	Update(id int, bookRequest m.BookUpdateRequest) (e.Book, error)
	Delete(id int) (e.Book, error)
}

type service struct {
	repository r.Repository
}

func NewService(repository r.Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]e.Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(id int) (e.Book, error) {
	return s.repository.FindById(id)
}

func (s *service) Create(bookRequest m.BookCreateRequest) (e.Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book := e.Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
		Discount:    int(discount),
	}

	return s.repository.Create(book)
}

func (s *service) Update(id int, bookRequest m.BookUpdateRequest) (e.Book, error) {
	book, _ := s.repository.FindById(id)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()

	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = int(rating)
	book.Discount = int(discount)

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(id int) (e.Book, error) {
	book, _ := s.repository.FindById(id)
	newBook, err := s.repository.Delete(book)
	return newBook, err
}
