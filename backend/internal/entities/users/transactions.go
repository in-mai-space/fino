package entities

import (
	"fino/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserTransactionInterface interface {
	InsertUser(user *models.User) (*models.User, error)
	RetrieveUser(id *uuid.UUID) (*models.User, error)
	UpdateUser(id *uuid.UUID, user *models.User) (*models.User, error)
	DeleteUser(id *uuid.UUID) error
}

type UserTransaction struct {
	db *gorm.DB
}

func NewUserTransaction(db *gorm.DB) *UserTransaction {
	return &UserTransaction{db: db}
}

func (transaction *UserTransaction) InsertUser(user *models.User) (*models.User, error) {
	if err := transaction.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (transaction *UserTransaction) RetrieveUser(id *uuid.UUID) (*models.User, error) {
	var user models.User
	if err := transaction.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (transaction *UserTransaction) UpdateUser(id *uuid.UUID, user *models.User) (*models.User, error) {
	if err := transaction.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (transaction *UserTransaction) DeleteUser(id *uuid.UUID) error {
	if err := transaction.db.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
