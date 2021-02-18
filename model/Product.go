package model

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	ID 		uint64			`gorm:"primary_key;auto_increment"`
	Sku		string			`gorm:"size:20;not null" json:"sku"`
	Name	string			`gorm:"size:40;not null" json:"name"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
