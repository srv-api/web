package entity

import "time"

type NewsComment struct {
	ID        string `gorm:"primaryKey"`
	BlogID    string `gorm:"index"`
	Name      string
	Email     string
	Comment   string
	CreatedAt time.Time
}
