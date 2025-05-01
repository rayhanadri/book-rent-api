package test

import (
	"library-api/model"
	"library-api/repository"
	"testing"

	"github.com/stretchr/testify/assert"

	"time"
)

func TestGetRentByID_Success(t *testing.T) {
	mockRepo := new(repository.MockRentRepository)

	// Representing a mock transaction object
	mockUserId := 1
	mockRentId := 1

	mockRent := model.Rent{
		ID:            mockRentId,
		BookID:        1,
		UserID:        mockUserId,
		Quantity:      1,
		TotalPrice:    5000,
		RentStartDate: time.Now(),
		RentEndDate:   time.Now().Add(7 * 24 * time.Hour),
		RentStatus:    "ACTIVE",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	mockRentPtr := &mockRent

	// Representing retrieving a transaction by ID from the database
	mockRepo.On("GetRentByID", mockUserId, mockRentId).Return(mockRentPtr, nil)
	rentPtr, err := mockRepo.GetRentByID(mockUserId, mockRentId)

	// Check if the transaction is retrieved successfully
	assert.NoError(t, err)
	assert.NotNil(t, rentPtr)

	mockRepo.AssertExpectations(t)
}

func TestGetRentByID_Failed(t *testing.T) {
	mockRepo := new(repository.MockRentRepository)

	// Representing retrieving a transaction by ID from the database
	mockUserId := 1
	mockRentId := 1
	mockRepo.On("GetRentByID", mockUserId, mockRentId).Return(nil, assert.AnError)
	rentPtr, err := mockRepo.GetRentByID(mockUserId, mockRentId)

	// Check if the transaction retrieval failed as expected
	assert.Error(t, err)
	assert.Nil(t, rentPtr)

	mockRepo.AssertExpectations(t)
}

func TestGetAllRent_Success(t *testing.T) {
	mockRepo := new(repository.MockRentRepository)

	// Representing a mock transaction object
	mockUserId := 1

	mockRents := []model.Rent{
		{
			ID:            1,
			BookID:        1,
			UserID:        mockUserId,
			Quantity:      1,
			TotalPrice:    5000,
			RentStartDate: time.Now(),
			RentEndDate:   time.Now().Add(7 * 24 * time.Hour),
			RentStatus:    "ACTIVE",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}
	mockRentsPtr := &mockRents

	// Representing retrieving a transaction by ID from the database
	mockRepo.On("GetAllRent", mockUserId).Return(mockRentsPtr, nil)
	rentsPtr, err := mockRepo.GetAllRent(mockUserId)

	// Check if the transaction is retrieved successfully
	assert.NoError(t, err)
	assert.NotNil(t, rentsPtr)

	mockRepo.AssertExpectations(t)
}

func TestGetAllRent_Failed(t *testing.T) {
	mockRepo := new(repository.MockRentRepository)

	// Representing retrieving a transaction by ID from the database
	mockUserId := 1
	mockRepo.On("GetAllRent", mockUserId).Return(nil, assert.AnError)
	rentsPtr, err := mockRepo.GetAllRent(mockUserId)

	// Check if the transaction retrieval failed as expected
	assert.Error(t, err)
	assert.Nil(t, rentsPtr)

	mockRepo.AssertExpectations(t)
}

func TestCreateRent_Success(t *testing.T) {
	mockRepo := new(repository.MockRentRepository)

	// Representing a mock transaction object
	mockUserId := 1

	mockRent := model.Rent{
		ID:            1,
		BookID:        1,
		UserID:        mockUserId,
		Quantity:      1,
		TotalPrice:    5000,
		RentStartDate: time.Now(),
		RentEndDate:   time.Now().Add(7 * 24 * time.Hour),
		RentStatus:    "ACTIVE",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	mockRentPtr := &mockRent

	// Representing retrieving a transaction by ID from the database
	mockRepo.On("CreateRent", mockUserId, mockRentPtr).Return(mockRentPtr, nil)
	rentPtr, err := mockRepo.CreateRent(mockUserId, mockRentPtr)

	// Check if the transaction is retrieved successfully
	assert.NoError(t, err)
	assert.NotNil(t, rentPtr)

	mockRepo.AssertExpectations(t)
}

func TestCreateRent_Failed(t *testing.T) {
	mockRepo := new(repository.MockRentRepository)

	// Representing retrieving a transaction by ID from the database
	mockUserId := 1
	mockRent := model.Rent{
		ID:            1,
		BookID:        1,
		UserID:        mockUserId,
		Quantity:      1,
		TotalPrice:    5000,
		RentStartDate: time.Now(),
		RentEndDate:   time.Now().Add(7 * 24 * time.Hour),
		RentStatus:    "ACTIVE",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	mockRentPtr := &mockRent

	mockRepo.On("CreateRent", mockUserId, mockRentPtr).Return(nil, assert.AnError)
	rentPtr, err := mockRepo.CreateRent(mockUserId, mockRentPtr)

	// Check if the transaction retrieval failed as expected
	assert.Error(t, err)
	assert.Nil(t, rentPtr)

	mockRepo.AssertExpectations(t)
}

func TestReturneRent_Success(t *testing.T) {
	mockRepo := new(repository.MockRentRepository)

	// Representing a mock transaction object
	// Representing a mock transaction object
	mockUserId := 1
	mockRentId := 1

	mockRent := model.Rent{
		ID:            1,
		BookID:        1,
		UserID:        mockUserId,
		Quantity:      1,
		TotalPrice:    5000,
		RentStartDate: time.Now(),
		RentEndDate:   time.Now().Add(7 * 24 * time.Hour),
		RentStatus:    "DONE",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	mockRentPtr := &mockRent

	// Representing retrieving a transaction by ID from the database
	mockRepo.On("ReturnRent", mockUserId, mockRentId).Return(mockRentPtr, nil)
	rentDone, err := mockRepo.ReturnRent(mockUserId, mockRentId)

	// Check if the transaction is retrieved successfully
	assert.NoError(t, err)
	assert.NotNil(t, rentDone)

	mockRepo.AssertExpectations(t)
}

func TestReturnRent_Failed(t *testing.T) {
	mockRepo := new(repository.MockRentRepository)

	// Representing retrieving a transaction by ID from the database
	mockUserId := 1
	mockRentId := 1
	mockRepo.On("ReturnRent", mockUserId, mockRentId).Return(nil, assert.AnError)
	rentDone, err := mockRepo.ReturnRent(mockUserId, mockRentId)

	// Check if the transaction retrieval failed as expected
	assert.Error(t, err)
	assert.Nil(t, rentDone)

	mockRepo.AssertExpectations(t)
}
