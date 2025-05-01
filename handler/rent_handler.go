package handler

import (
	"fmt"
	"library-api/config"
	"library-api/model"
	"library-api/repository"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type RentHandler interface {
	GetRentByID(c echo.Context) error
	GetAllRent(c echo.Context) error
	CreateRent(c echo.Context) error
	ReturnRent(c echo.Context) error
}

type rentHandler struct {
	rentRepo repository.RentRepository
}

func NewRentHandler(rentRepo repository.RentRepository) RentHandler {
	return &rentHandler{rentRepo: rentRepo}
}

// GetRentsByID godoc
// @Summary Get a rent by ID
// @Description Retrieve a rent's details by its ID
// @Tags rents
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>"
// @Param id path int true "Rent ID" // Updated parameter description
// @Success 200 {object} model.Response
// @Router /rents/{id} [get]
func (h *rentHandler) GetRentByID(c echo.Context) error {
	//get user id from context
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Status:  http.StatusUnauthorized,
			Message: "User not authenticated",
		})
	}

	userIdFloat, ok := userID.(float64)
	if !ok {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid user ID",
		})
	}
	userIdInt := int(userIdFloat)

	//get rent id from param
	rentID := c.Param("id")
	if rentID == "" {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Rent ID is required",
		})
	}

	// Convert rentID to int
	rentIDInt, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid rent ID",
		})
	}

	// get rent by id
	rent, err := h.rentRepo.GetRentByID(userIdInt, rentIDInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to retrieve rent",
		})
	}

	// return rent
	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    rent,
	})
}

// GetAllRent godoc
// @Summary Get all rents
// @Description Retrieve all rents for a user
// @Tags rents
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>"
// @Success 200 {object} model.Response
// @Router /rents [get] // Updated the router path to get all rents
func (h *rentHandler) GetAllRent(c echo.Context) error {
	//get user id from context
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Status:  http.StatusUnauthorized,
			Message: "User not authenticated",
		})
	}

	userIdFloat, ok := userID.(float64)
	if !ok {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid user ID",
		})
	}
	userIdInt := int(userIdFloat)

	// get all rent
	rents, err := h.rentRepo.GetAllRent(userIdInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// check if user own the rent
	for _, rent := range *rents {
		if rent.UserID != userIdInt {
			return c.JSON(http.StatusForbidden, model.Response{
				Status:  http.StatusForbidden,
				Message: "You are not allowed to access this rent data",
			})
		}
	}

	// return all rent
	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    rents,
	})
}

// Create Rent godoc
// @Summary Create a new rent
// @Description Create a new rent for a user
// @Tags rents
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>"
// @Success 201 {object} model.Response
// @Router /rents [post] // Updated the router path to use POST method
func (h *rentHandler) CreateRent(c echo.Context) error {
	//get user id from context
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Status:  http.StatusUnauthorized,
			Message: "User not authenticated",
		})
	}

	userIdFloat, ok := userID.(float64)
	if !ok {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid user ID",
		})
	}
	userIdInt := int(userIdFloat)

	// bind request to rent model
	rent := new(model.Rent)
	if err := c.Bind(&rent); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
		})
	}

	fmt.Println(rent)

	// set user id to rent
	rent.UserID = userIdInt
	rent.RentStatus = "PENDING"

	// get book id from rent
	var book model.Book
	var db = config.DB

	err := db.Where("id = ?", rent.BookID).First(&book).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  http.StatusNotFound,
			Message: "Book not found",
		})
	}

	// fmt.Println(book)

	// check rent qty
	if rent.Quantity > book.Stock {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Rent quantity exceeds book stock",
		})
	}

	// update book stock
	book.Stock -= rent.Quantity
	if book.Stock == 0 {
		book.Available = false
	} else {
		book.Available = true
	}

	// update book
	err = db.Save(&book).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update book data",
		})
	}

	// create rent
	rent, err = h.rentRepo.CreateRent(userIdInt, rent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create rent",
		})
	}

	// return rent
	return c.JSON(http.StatusCreated, model.Response{
		Status:  http.StatusCreated,
		Message: "Rent created successfully",
		Data:    rent,
	})
}

// ReturnRent godoc
// @Summary Return a rent
// @Description Process the return of a rent by its ID
// @Tags rents
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <access_token>"
// @Param id path int true "Rent ID"
// @Success 200 {object} model.Response
// @Router /rents/return/{id} [put]
func (h *rentHandler) ReturnRent(c echo.Context) error {
	//get user id from context
	userID := c.Get("user_id")
	if userID == nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Status:  http.StatusUnauthorized,
			Message: "User not authenticated",
		})
	}

	userIdFloat, ok := userID.(float64)
	if !ok {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid user ID",
		})
	}
	userIdInt := int(userIdFloat)

	//get rent id from param
	rentID := c.Param("id")
	if rentID == "" {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Rent ID is required",
		})
	}

	// Convert rentID to int
	rentIDInt, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid rent ID",
		})
	}

	//Check existing rent by ID
	existingRent, err := h.rentRepo.GetRentByID(userIdInt, rentIDInt)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  http.StatusNotFound,
			Message: "Rent not found",
		})
	}

	if existingRent.RentStatus == "DONE" {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Rent already returned",
		})
	}

	rent := new(model.Rent)
	rent, err = h.rentRepo.ReturnRent(userIdInt, rentIDInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  http.StatusInternalServerError,
			Message: "Failed to return rent",
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: "Rent returned successfully",
		Data:    rent,
	})
}
