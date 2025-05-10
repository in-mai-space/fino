package models

import (
	"time"

	"github.com/google/uuid"
)

type IntegrationNotion struct {
	UserID              uint       `gorm:"primaryKey"`
	AccessToken         string     `gorm:"type:varchar(255);not null"`
	RefreshToken        string     `gorm:"type:varchar(255);not null"`
	NotionUserID        string     `gorm:"type:varchar(100);not null"`
	NotionWorkspaceID   string     `gorm:"type:varchar(100);not null"`
	NotionWorkspaceName string     `gorm:"type:varchar(255);not null"`
	TokenExpiresAt      *time.Time `gorm:"type:timestamp"`
	CreatedAt           time.Time  `gorm:"default:current_timestamp"`
	UpdatedAt           time.Time  `gorm:"default:current_timestamp"`

	User            User             `gorm:"foreignKey:UserID"`
	NotionDatabases []NotionDatabase `gorm:"foreignKey:IntegrationUserID;constraint:OnDelete:CASCADE"`
}

type NotionDatabase struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	IntegrationUserID uint      `gorm:"not null;index"`
	NotionDatabaseID  string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Title             string    `gorm:"type:varchar(255);not null"`
	Schema            string    `gorm:"type:text"`
	IsActive          bool      `gorm:"default:true"`
	CreatedAt         time.Time `gorm:"default:current_timestamp"`
	UpdatedAt         time.Time `gorm:"default:current_timestamp"`

	Integration     IntegrationNotion `gorm:"foreignKey:IntegrationUserID"`
	AccountMappings []AccountMapping  `gorm:"foreignKey:NotionDatabaseID"`
}
