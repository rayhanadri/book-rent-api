package repository

import (
	"errors"
	"library-api/model"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAllTransaction(user_id int) (*[]model.Transaction, error)
	CreateTransaction(user_id int, transaction *model.Transaction) (*model.Transaction, error)
	GetTransactionByID(user_id int, transactionID int) (*model.Transaction, error)
	UpdateTransaction(user_id int, transaction *model.Transaction) (*model.Transaction, error)
	CancelTransaction(user_id int, transactionID int) (*model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) GetAllTransaction(user_id int) (*[]model.Transaction, error) {
	// validate user id
	user := new(model.User)
	if err := r.db.Where("id = ?", user_id).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	var transactions *[]model.Transaction
	if err := r.db.Where("user_id = ?", user_id).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *transactionRepository) CreateTransaction(user_id int, transaction *model.Transaction) (*model.Transaction, error) {
	// validate user id
	user := new(model.User)
	if err := r.db.Where("id = ?", user_id).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	transaction.UserID = user_id

	// validate transaction data
	if transaction.TransactionType == "" || transaction.Amount <= 0 || transaction.PaymentMethod == "" || transaction.Description == "" {
		return nil, errors.New("transaction type, amount, payment method, and description are required")
	}
	if user_id <= 0 {
		return nil, errors.New("user_id must be valid")
	}

	if transaction.TransactionType == "Rent" && transaction.RentID <= 0 {
		return nil, errors.New("rent_id must be valid")
	}
	//

	if transaction.TransactionType == "Topup" {
		if err := r.db.Omit("rent_id").Create(transaction).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.db.Create(transaction).Error; err != nil {
			return nil, err
		}
	}

	return transaction, nil
}

func (r *transactionRepository) UpdateTransaction(user_id int, transaction *model.Transaction) (*model.Transaction, error) {
	// validate user id
	user := new(model.User)
	if err := r.db.Where("id = ?", user_id).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	transaction.UserID = user_id

	if transaction.TransactionType != "Rent" {
		if err := r.db.Omit("rent_id").Save(transaction).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.db.Save(transaction).Error; err != nil {
			return nil, err
		}
	}

	return transaction, nil
}

func (r *transactionRepository) CancelTransaction(user_id int, transactionID int) (*model.Transaction, error) {
	// validate user id
	user := new(model.User)
	if err := r.db.Where("id = ?", user_id).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// check if transaction exists
	transaction := new(model.Transaction)
	if err := r.db.Where("user_id = ? AND id = ?", user_id, transactionID).First(transaction).Error; err != nil {
		return nil, errors.New("transaction not found")
	}

	transaction.Status = "CANCELED"
	transaction.UpdatedAt = time.Now()

	if transaction.TransactionType != "Rent" {
		if err := r.db.Omit("rent_id").Save(transaction).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.db.Save(transaction).Error; err != nil {
			return nil, err
		}
	}

	return transaction, nil
}

func (r *transactionRepository) GetTransactionByID(user_id int, transactionID int) (*model.Transaction, error) {
	// validate user id
	user := new(model.User)
	if err := r.db.Where("id = ?", user_id).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	transaction := new(model.Transaction)
	if err := r.db.Where("user_id = ? AND id = ?", user_id, transactionID).First(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}
