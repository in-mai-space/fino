package models

type AccountMapping struct {
	ID               uint           `gorm:"primaryKey"`
	PlaidAccountID   uint           `gorm:"not null"`
	NotionDatabaseID uint           `gorm:"not null"`
	PlaidAccount     PlaidAccount   `gorm:"foreignKey:PlaidAccountID"`
	NotionDatabase   NotionDatabase `gorm:"foreignKey:NotionDatabaseID"`
}

