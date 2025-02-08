package models

import "time"

type RefreshToken struct {
	ID        uint   `gorm:"primaryKey"`
	TokenHash string `gorm:"unique"`
	UserID    uint
	ExpiresAt time.Time
	Revoked   bool
	CreatedAt time.Time
}
