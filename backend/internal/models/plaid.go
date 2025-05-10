package models

import (
	"time"

	"github.com/google/uuid"
)

type IntegrationPlaid struct {
	UserID          uint   `gorm:"primaryKey"`
	AccessToken     string `gorm:"size:255;not null"`
	ItemID          string `gorm:"size:255;not null;index"`
	InstitutionID   string `gorm:"size:100;not null"`
	InstitutionName string `gorm:"size:255;not null"`
	WebhookURL      string `gorm:"size:255"`
	WebhookEnabled  bool   `gorm:"default:false"`
	LastFetchStatus string `gorm:"size:50"`
	LastFetchTime   time.Time
	CreatedAt       time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	User            User           `gorm:"foreignKey:UserID"`
	PlaidAccounts   []PlaidAccount `gorm:"foreignKey:IntegrationUserID;constraint:OnDelete:CASCADE"`
}

type PlaidAccount struct {
	ID                uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	IntegrationUserID uint             `gorm:"not null"`
	PlaidAccountID    string           `gorm:"size:100;not null;uniqueIndex"`
	Name              string           `gorm:"size:255;not null"`
	Mask              string           `gorm:"size:10"`
	Type              string           `gorm:"size:50;not null"`
	Subtype           string           `gorm:"size:50"`
	IsActive          bool             `gorm:"default:true"`
	CreatedAt         time.Time        `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time        `gorm:"default:CURRENT_TIMESTAMP"`
	Integration       IntegrationPlaid `gorm:"foreignKey:IntegrationUserID"`
	Transactions      []Transaction    `gorm:"foreignKey:PlaidAccountID"`
	AccountMappings   []AccountMapping `gorm:"foreignKey:PlaidAccountID"`
}
