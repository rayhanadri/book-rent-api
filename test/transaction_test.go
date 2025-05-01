package test

import (
	"library-api/model"
	"library-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"

	"time"
)

func TestGetTransactionByID_Success(t *testing.T) {
	mockRepo := new(repository.MockTransactionRepository)

	// Representing a mock transaction object
	mockUserId := 1
	mockTransaction := model.Transaction{
		ID:              1,
		TransactionType: "Rent",
		PaymentMethod:   "Payment Gateway",
		Amount:          5000,
		Status:          "PENDING",
		Description:     "Rent Book",
		UserID:          mockUserId,
		InvoiceID:       "INV-12345",
		InvoiceURL:      "https://example.com/invoice/12345",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		RentID:          1,
	}
	mockTransactionPtr := &mockTransaction

	// Representing retrieving a transaction by ID from the database
	mockRepo.On("GetTransactionByID", mockUserId, 1).Return(mockTransactionPtr, nil)
	transactionPtr, err := mockRepo.GetTransactionByID(mockUserId, 1)

	// Check if the transaction is retrieved successfully
	assert.NoError(t, err)
	assert.NotNil(t, transactionPtr)

	mockRepo.AssertExpectations(t)
}

func TestGetTransactionByID_Failed(t *testing.T) {
	mockRepo := new(repository.MockTransactionRepository)

	// Representing retrieving a transaction by ID from the database
	mockUserId := 1
	mockRepo.On("GetTransactionByID", mockUserId, 1).Return(nil, assert.AnError)
	transactionPtr, err := mockRepo.GetTransactionByID(mockUserId, 1)

	// Check if the transaction retrieval failed as expected
	assert.Error(t, err)
	assert.Nil(t, transactionPtr)

	mockRepo.AssertExpectations(t)
}

func TestGetAllTransaction_Success(t *testing.T) {
	mockRepo := new(repository.MockTransactionRepository)

	// Representing transactions retrieved from the database
	mockUserId := 1
	mockTransactions := []model.Transaction{
		{
			ID:              1,
			TransactionType: "Rent",
			PaymentMethod:   "Payment Gateway",
			Amount:          5000,
			Status:          "PENDING",
			Description:     "Rent Book",
			UserID:          mockUserId,
			InvoiceID:       "INV-12345",
			InvoiceURL:      "https://example.com/invoice/12345",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			RentID:          1,
		},
		{
			ID:              2,
			TransactionType: "Topup",
			PaymentMethod:   "Payment Gateway",
			Amount:          50000,
			Status:          "PENDING",
			Description:     "Topup balance",
			UserID:          mockUserId,
			InvoiceID:       "INV-12346",
			InvoiceURL:      "https://example.com/invoice/12346",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}

	mockTransactionsPtr := &mockTransactions

	// Representing retrieving all transactions from the database
	mockRepo.On("GetAllTransaction", mockUserId).Return(mockTransactionsPtr, nil)
	transactions, err := mockRepo.GetAllTransaction(mockUserId)

	// Check if the transaction is retrieved successfully
	assert.NoError(t, err)
	assert.NotNil(t, transactions)

	mockRepo.AssertExpectations(t)
}

func TestGetAllTransaction_Failed(t *testing.T) {
	mockRepo := new(repository.MockTransactionRepository)

	// Representing retrieving all transactions from the database
	mockUserId := 1
	mockRepo.On("GetAllTransaction", mockUserId).Return(nil, assert.AnError)
	transactions, err := mockRepo.GetAllTransaction(mockUserId)

	// Check if the transaction retrieval failed as expected
	assert.Error(t, err)
	assert.Nil(t, transactions)

	mockRepo.AssertExpectations(t)
}

func TestCreateTransaction_Success(t *testing.T) {
	mockRepo := new(repository.MockTransactionRepository)

	// Representing a transaction created and retrieved from the database
	mockUserId := 1
	mockTransaction := model.Transaction{
		ID:              1,
		TransactionType: "Rent",
		PaymentMethod:   "Payment Gateway",
		Amount:          5000,
		Status:          "PENDING",
		Description:     "Rent Book",
		UserID:          mockUserId,
		InvoiceID:       "INV-12345",
		InvoiceURL:      "https://example.com/invoice/12345",
		RentID:          1,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mockTransactionPtr := &mockTransaction

	// Representing creating a transaction in the database
	mockRepo.On("CreateTransaction", mockUserId, mockTransactionPtr).Return(mockTransactionPtr, nil)
	transactionPtr, err := mockRepo.CreateTransaction(mockUserId, mockTransactionPtr)

	// Check if the transaction is created successfully
	assert.NoError(t, err)
	assert.NotNil(t, transactionPtr)

	mockRepo.AssertExpectations(t)
}

func TestCreateTransaction_Failed(t *testing.T) {
	mockRepo := new(repository.MockTransactionRepository)

	// Representing creating a transaction in the database
	mockUserId := 1
	mockTransaction := model.Transaction{
		ID:              1,
		TransactionType: "Rent",
		PaymentMethod:   "Payment Gateway",
		Amount:          5000,
		Status:          "PENDING",
		Description:     "Rent Book",
		UserID:          mockUserId,
		InvoiceID:       "INV-12345",
		InvoiceURL:      "https://example.com/invoice/12345",
		RentID:          1,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mockTransactionPtr := &mockTransaction

	mockRepo.On("CreateTransaction", mockUserId, mockTransactionPtr).Return(nil, assert.AnError)
	transactionPtr, err := mockRepo.CreateTransaction(mockUserId, mockTransactionPtr)

	// Check if the transaction creation failed as expected
	assert.Error(t, err)
	assert.Nil(t, transactionPtr)

	mockRepo.AssertExpectations(t)
}

func TestUpdateTransaction_Success(t *testing.T) {
	mockRepo := new(repository.MockTransactionRepository)

	// Representing a transaction created and retrieved from the database
	mockUserId := 1
	mockTransaction := model.Transaction{
		ID:              1,
		TransactionType: "Rent",
		PaymentMethod:   "Payment Gateway",
		Amount:          5000,
		Status:          "PENDING",
		Description:     "Rent Book",
		UserID:          mockUserId,
		InvoiceID:       "INV-12345",
		InvoiceURL:      "https://example.com/invoice/12345",
		RentID:          1,
	}
	mockTransactionPtr := &mockTransaction

	// Representing updating a transaction in the database
	mockRepo.On("UpdateTransaction", mockUserId, mockTransactionPtr).Return(mockTransactionPtr, nil)
	transactionPtr, err := mockRepo.UpdateTransaction(mockUserId, mockTransactionPtr)

	// Check if the transaction is updated successfully
	assert.NoError(t, err)
	assert.NotNil(t, transactionPtr)

	mockRepo.AssertExpectations(t)
}

func TestUpdateTransaction_Failed(t *testing.T) {
	mockRepo := new(repository.MockTransactionRepository)

	// Representing a transaction created and retrieved from the database
	mockUserId := 1
	mockTransaction := model.Transaction{
		ID:              1,
		TransactionType: "Rent",
		PaymentMethod:   "Payment Gateway",
		Amount:          5000,
		Status:          "PENDING",
		Description:     "Rent Book",
		UserID:          mockUserId,
		InvoiceID:       "INV-12345",
		InvoiceURL:      "https://example.com/invoice/12345",
		RentID:          1,
	}
	mockTransactionPtr := &mockTransaction

	// Representing updating a transaction in the database
	mockRepo.On("UpdateTransaction", mockUserId, mockTransactionPtr).Return(nil, assert.AnError)
	transactionPtr, err := mockRepo.UpdateTransaction(mockUserId, mockTransactionPtr)

	// Check if the transaction update failed as expected
	assert.Error(t, err)
	assert.Nil(t, transactionPtr)

	mockRepo.AssertExpectations(t)
}
