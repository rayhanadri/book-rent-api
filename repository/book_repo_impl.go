package repository

import (
	"library-api/model"

	"github.com/stretchr/testify/mock"
)

type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) GetBooksByID(id int) (*model.Book, error) {
	args := m.Called(id)
	if book := args.Get(0); book != nil {
		return book.(*model.Book), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockBookRepository) GetAllBooks() ([]*model.Book, error) {
	args := m.Called()
	if books := args.Get(0); books != nil {
		return books.([]*model.Book), args.Error(1)
	}
	return nil, args.Error(1)
}
