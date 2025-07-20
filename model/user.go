package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Mobile    string `gorm:"uniqueIndex"`
	CreatedAt time.Time
}
