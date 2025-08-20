package models

import (
	"time"
)

type BlacklistedToken struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    JTI       string   `json:"jti" gorm:"uniqueIndex"`
    ExpiresAt time.Time `gorm:"uniqueIndex"`
}

// JTI - a unique identifier for the JWT this can be used to prevent JWT token abuse and to ensure uniqueness of token