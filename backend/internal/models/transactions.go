package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID             uint      `gorm:"not null"`
	PlaidAccountID     uint      `gorm:"not null"`
	PlaidTransactionID string    `gorm:"size:100;uniqueIndex;not null"`
	Date               time.Time `gorm:"not null"`
	Name               string    `gorm:"size:255;not null"`
	Amount             float64   `gorm:"not null"`
	Category           string    `gorm:"size:255"`
	Pending            bool      `gorm:"default:false"`
	NotionPageID       string    `gorm:"size:100"`
	LastSyncedAt       time.Time
	SyncStatus         SyncStatusEnum `gorm:"default:'pending'"`
	SyncError          string         `gorm:"size:255"`
	CreatedAt          time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	User               User           `gorm:"foreignKey:UserID"`
	PlaidAccount       PlaidAccount   `gorm:"foreignKey:PlaidAccountID"`
}
