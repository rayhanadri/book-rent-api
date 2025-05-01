package repository

import (
	"errors"
	"library-api/model"

	"gorm.io/gorm"
)

type RentRepository interface {
	GetAllRent(user_id int) (*[]model.Rent, error)
	GetRentByID(user_id int, id int) (*model.Rent, error)
	CreateRent(user_id int, rent *model.Rent) (*model.Rent, error)
	ReturnRent(user_id int, id int) (*model.Rent, error)
}

type rentRepository struct {
	db *gorm.DB
}

func NewRentRepository(db *gorm.DB) RentRepository {
	return &rentRepository{db: db}
}

func (r *rentRepository) GetRentByID(user_id int, id int) (*model.Rent, error) {
	var rent model.Rent
	if err := r.db.Where("id = ? AND user_id = ?", id, user_id).First(&rent).Error; err != nil {
		return nil, err
	}
	return &rent, nil
}

func (r *rentRepository) GetAllRent(user_id int) (*[]model.Rent, error) {
	var rents []model.Rent
	if err := r.db.Where("user_id = ?", user_id).Find(&rents).Error; err != nil {
		return nil, err
	}
	return &rents, nil
}

func (r *rentRepository) CreateRent(user_id int, rent *model.Rent) (*model.Rent, error) {
	// set rent user id
	rent.UserID = user_id

	// validate rent data
	if rent.BookID <= 0 || rent.Quantity <= 0 || rent.TotalPrice <= 0 {
		return nil, errors.New("book_id, quantity, and total_price are required")
	}
	if user_id <= 0 {
		return nil, errors.New("user_id must be valid")
	}
	if rent.RentStatus == "" {
		return nil, errors.New("rent status is required")
	}

	if rent.RentStartDate.IsZero() {
		return nil, errors.New("rent start date is required")
	}
	if rent.RentEndDate.IsZero() {
		return nil, errors.New("rent end date is required")
	}
	if rent.RentEndDate.Before(rent.RentStartDate) {
		return nil, errors.New("rent end date must be after rent start date")
	}
	//

	// create rent data
	err := r.db.Omit("id").Create(&rent).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Last(&rent).Error
	if err != nil {
		return nil, err
	}

	return rent, nil
}

func (r *rentRepository) ReturnRent(user_id int, id int) (*model.Rent, error) {
	//get rent data by id
	var rent model.Rent

	//find rent by id and user id
	err := r.db.Where("id = ? AND user_id = ?", id, user_id).First(&rent).Error
	if err != nil {
		return nil, errors.New("rent not found")
	}

	//update rent status to DONE and save it
	rent.RentStatus = "DONE"
	err = r.db.Save(&rent).Error
	if err != nil {
		return nil, err
	}

	// Update the book stock
	var book model.Book
	err = r.db.Where("id = ?", rent.BookID).First(&book).Error
	if err != nil {
		return nil, errors.New("book not found")
	}

	// Update quantity and availability of the book
	book.Stock += rent.Quantity
	if book.Stock >= 1 {
		book.Available = true
	} else {
		book.Available = false
	}

	// Save the updated book data
	err = r.db.Save(&book).Error
	if err != nil {
		return nil, errors.New("failed to update book stock")
	}
	//

	return &rent, nil
}
