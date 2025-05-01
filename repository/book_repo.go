package repository

import (
	"library-api/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetBooksByID(id int) (*model.Book, error)
	GetAllBooks() ([]*model.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetBooksByID(id int) (*model.Book, error) {
	var book model.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) GetAllBooks() ([]*model.Book, error) {
	var books []*model.Book
	if err := r.db.Where("stock > 0").Order("id ASC").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
