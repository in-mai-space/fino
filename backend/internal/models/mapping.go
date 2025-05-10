package models

import "github.com/google/uuid"

type AccountMapping struct {
	ID               uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PlaidAccountID   uint           `gorm:"not null"`
	NotionDatabaseID uint           `gorm:"not null"`
	PlaidAccount     PlaidAccount   `gorm:"foreignKey:PlaidAccountID"`
	NotionDatabase   NotionDatabase `gorm:"foreignKey:NotionDatabaseID"`
}
