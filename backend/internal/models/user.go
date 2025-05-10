package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                 uuid.UUID           `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
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
