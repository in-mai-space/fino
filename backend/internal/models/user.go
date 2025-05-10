package models

import (
	"time"
)

type User struct {
	ID                 uint                `gorm:"primaryKey"`
	FirstName          string              `gorm:"size:100;not null"`
	LastName           string              `gorm:"size:100;not null"`
	Email              string              `gorm:"size:255;uniqueIndex;not null"`
	CreatedAt          time.Time           `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time           `gorm:"default:CURRENT_TIMESTAMP"`
	Transactions       []Transaction       `gorm:"foreignKey:UserID"`
	SyncConfigurations []SyncConfiguration `gorm:"foreignKey:UserID"`
	SyncLogs           []SyncLog           `gorm:"foreignKey:UserID"`
	NotionIntegration  *IntegrationNotion  `gorm:"foreignKey:UserID"`
	PlaidIntegration   *IntegrationPlaid   `gorm:"foreignKey:UserID"`
}
