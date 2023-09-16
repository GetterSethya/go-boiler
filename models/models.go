package models

import (
	"time"

	"gorm.io/gorm"
)

type Grocery struct {
	gorm.Model
	Name string `json:"name"`
	Quantity int `json:"quantity"`
}


type Model struct {
	Id uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt gorm.DeletedAt `gorm:"index"`
}