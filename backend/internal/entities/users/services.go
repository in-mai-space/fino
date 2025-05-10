package entities

import (
	"fino/internal/models"
	"fino/internal/utilities"
	"fmt"

	"github.com/google/uuid"
)

type UserServiceInterface interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(id *uuid.UUID) (*models.User, error)
	UpdateUser(id *uuid.UUID, user *models.User) (*models.User, error)
	DeleteUser(id *uuid.UUID) error
}

type UserService struct {
	tx UserTransactionInterface
}

func NewUserService(tx UserTransactionInterface) *UserService {
	return &UserService{tx: tx}
}

func (us *UserService) CreateUser(user *models.User) (*models.User, error) {
	user, err := us.tx.InsertUser(user)
	if err != nil {
		return nil, utilities.NewInternalServerError(fmt.Sprintf("Failed to create user: %v", err.Error()))
	}
	return user, nil
}

func (us *UserService) GetUser(id *uuid.UUID) (*models.User, error) {
	user, err := us.tx.RetrieveUser(id)
	if err != nil {
		return nil, utilities.NewNotFoundError(fmt.Sprintf("User does not exist: %v", err.Error()))
	}
	return user, nil
}

func (us *UserService) UpdateUser(id *uuid.UUID, user *models.User) (*models.User, error) {
	user, err := us.tx.UpdateUser(id, user)
	if err != nil {
		return nil, utilities.NewInternalServerError(fmt.Sprintf("Failed to update user: %v", err.Error()))
	}
	return user, nil
}

func (us *UserService) DeleteUser(id *uuid.UUID) error {
	err := us.tx.DeleteUser(id)
	if err != nil {
		return utilities.NewInternalServerError(fmt.Sprintf("Failed to delete user: %v", err.Error()))
	}
	return nil
}
