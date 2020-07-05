package models

import (
	"time"
)

type Order struct {
	ID              uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserID          uint32    `gorm:"not null" json:"user_id"`
	ItemDescription string    `gorm:"size:255;not null" json:"item_description"`
	ItemQuantity    uint32    `gorm:"not null" json:"item_quantity"`
	ItemPrice       uint32    `gorm:"not null" json:"item_price"`
	TotalValue      uint32    `gorm:"not null" json:"total_value"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
