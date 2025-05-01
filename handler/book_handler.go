package handler

import (
	"library-api/model"
	"library-api/repository"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type BookHandler interface {
	GetBooksByID(c echo.Context) error
	GetAllBooks(c echo.Context) error
}

type bookHandler struct {
	bookRepo repository.BookRepository
}

func NewBookHandler(bookRepo repository.BookRepository) BookHandler {
	return &bookHandler{bookRepo: bookRepo}
}

// GetBooksByID godoc
// @Summary Get a book by ID
// @Description Retrieve a book's details by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} model.Response
// @Router /books/{id} [get]
func (h *bookHandler) GetBooksByID(c echo.Context) error {
	bookID := c.Param("id")
	if bookID == "" {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Book ID is required",
		})
	}

	// Convert bookID to int
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid book ID",
		})
	}

	book, err := h.bookRepo.GetBooksByID(bookIDInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    book,
	})
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Retrieve all books' details
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /books [get] // Updated the router path to retrieve all books
func (h *bookHandler) GetAllBooks(c echo.Context) error {
	books, err := h.bookRepo.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    books,
	})
}
