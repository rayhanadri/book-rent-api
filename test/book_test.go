package test

import (
	"library-api/model"
	"library-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"

	"time"
)

func TestGetBooksByID_Success(t *testing.T) {
	mockRepo := new(repository.MockBookRepository)

	// Representing a book retrieved from the database
	mockBook := model.Book{
		ID:          1,
		Title:       "Book Title",
		Author:      "Author Name",
		Publisher:   "Publisher Name",
		PublishedAt: time.Now(),
		ISBN:        "1234567890",
		Category:    "Fiction",
		Stock:       2,
		Available:   true,
		Price:       5000,
		Description: "Book Description",
	}
	mockBookPtr := &mockBook

	// Representing retrieving a book by ID from the database
	mockRepo.On("GetBooksByID", 1).Return(mockBookPtr, nil)
	bookPtr, err := mockRepo.GetBooksByID(1)

	// Check if the book is retrieved successfully
	assert.NoError(t, err)
	assert.Equal(t, mockBookPtr, bookPtr)

	mockRepo.AssertExpectations(t)
}

func TestGetBooksByID_Failed(t *testing.T) {
	mockRepo := new(repository.MockBookRepository)

	// Representing retrieving a book by ID from the database
	mockRepo.On("GetBooksByID", 2).Return(nil, assert.AnError)
	bookPtr, err := mockRepo.GetBooksByID(2)

	// Check if the book retrieval failed as expected
	assert.Error(t, err)
	assert.Nil(t, bookPtr)

	mockRepo.AssertExpectations(t)
}

func TestGetAllBooks_Success(t *testing.T) {
	mockRepo := new(repository.MockBookRepository)

	// Representing a list of books retrieved from the database
	mockBooks := []*model.Book{
		{
			ID:          1,
			Title:       "Book Title 1",
			Author:      "Author Name 1",
			Publisher:   "Publisher Name 1",
			PublishedAt: time.Now(),
			ISBN:        "1234567890",
			Category:    "Fiction",
			Stock:       5,
			Available:   true,
			Price:       4500,
			Description: "Book Description 1",
		},
		{
			ID:          2,
			Title:       "Book Title 2",
			Author:      "Author Name 2",
			Publisher:   "Publisher Name 2",
			PublishedAt: time.Now(),
			ISBN:        "0987654321",
			Category:    "Non-Fiction",
			Stock:       3,
			Available:   true,
			Price:       6000,
			Description: "Book Description 2",
		},
	}

	// Representing retrieving all books from the database
	mockRepo.On("GetAllBooks").Return(mockBooks, nil)
	books, err := mockRepo.GetAllBooks()

	// Check if all books are retrieved successfully
	assert.NoError(t, err)
	assert.Equal(t, mockBooks, books)

	mockRepo.AssertExpectations(t)
}
