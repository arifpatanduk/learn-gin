package models

import (
	"time"
)

type Car struct {
	ID uint `gorm:"primaryKey"`
	Brand string `gorm:"not null;type:varchar(191)"`
	Model string `gorm:"not null;type:varchar(191)"`
	Price int `gorm:"not null;type:int"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
}