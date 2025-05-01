package repository

import (
	"library-api/model"

	"github.com/stretchr/testify/mock"
)

type MockRentRepository struct {
	mock.Mock
}

func (m *MockRentRepository) GetRentByID(user_id int, id int) (*model.Rent, error) {
	args := m.Called(user_id, id)
	if rent := args.Get(0); rent != nil {
		return rent.(*model.Rent), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRentRepository) GetAllRent(user_id int) (*[]model.Rent, error) {
	args := m.Called(user_id)
	if rents := args.Get(0); rents != nil {
		return rents.(*[]model.Rent), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRentRepository) CreateRent(user_id int, rent *model.Rent) (*model.Rent, error) {
	args := m.Called(user_id, rent)
	if rent := args.Get(0); rent != nil {
		return rent.(*model.Rent), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRentRepository) ReturnRent(user_id int, id int) (*model.Rent, error) {
	args := m.Called(user_id, id)
	if rent := args.Get(0); rent != nil {
		return rent.(*model.Rent), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRentRepository) CancelRent(user_id int, id int) (*model.Rent, error) {
	args := m.Called(user_id, id)
	if rent := args.Get(0); rent != nil {
		return rent.(*model.Rent), args.Error(1)
	}
	return nil, args.Error(1)
}
