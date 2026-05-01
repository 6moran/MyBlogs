package entity

import "time"

type Message struct {
	ID         int    `gorm:"primaryKey"`
	FromUserID int    `gorm:"not null;default:0"`
	UserID     int    `gorm:"not null"`
	Title      string `gorm:"size:255;not null"`
	Content    string `gorm:"size:500;not null"`
	IsRead     int    `gorm:"default:0"`
	CreatedAt  time.Time
}
