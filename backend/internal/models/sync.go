package models

import (
	"time"
)

type SyncInterval string

const (
	HOURLY  SyncInterval = "hourly"
	DAILY   SyncInterval = "daily"
	WEEKLY  SyncInterval = "weekly"
	MONTHLY SyncInterval = "monthly"
	MANUAL  SyncInterval = "manual"
)

type SyncStatusEnum string

const (
	PENDING     SyncStatusEnum = "pending"
	IN_PROGRESS SyncStatusEnum = "in_progress"
	COMPLETED   SyncStatusEnum = "completed"
	FAILED      SyncStatusEnum = "failed"
	CANCELLED   SyncStatusEnum = "cancelled"
)

type TriggerTypeEnum string

const (
	WEBHOOK        TriggerTypeEnum = "webhook"
	SCHEDULED      TriggerTypeEnum = "scheduled"
	MANUAL_TRIGGER TriggerTypeEnum = "manual"
	INITIAL_SETUP  TriggerTypeEnum = "initial_setup"
	RETRY          TriggerTypeEnum = "retry"
)

type SyncConfiguration struct {
	ID                 uint         `gorm:"primaryKey"`
	UserID             uint         `gorm:"not null"`
	SyncInterval       SyncInterval `gorm:"default:'daily'"`
	SyncEnabled        bool         `gorm:"default:true"`
	WebhookSyncEnabled bool         `gorm:"default:true"`
	IncludePending     bool         `gorm:"default:true"`
	MinTransactionDate time.Time
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	User               User      `gorm:"foreignKey:UserID"`
}

type SyncLog struct {
	ID           uint            `gorm:"primaryKey"`
	UserID       uint            `gorm:"not null"`
	Status       SyncStatusEnum  `gorm:"default:'pending'"`
	TriggerType  TriggerTypeEnum `gorm:"not null"`
	ErrorMessage string          `gorm:"size:255"`
	WorkflowID   string          `gorm:"size:100"`
	CreatedAt    time.Time       `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time       `gorm:"default:CURRENT_TIMESTAMP"`
	User         User            `gorm:"foreignKey:UserID"`
}
