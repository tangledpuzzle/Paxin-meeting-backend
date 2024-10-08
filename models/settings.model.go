package models

import (
	"time"
)

type Langs struct {
	ID        uint       `gorm:"primary_key"`
	Name      string     `gorm:"not null"`
	Code      string     `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}
