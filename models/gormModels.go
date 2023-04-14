package models

import (
	"time"
)



type GormModel struct {
// ID uint `gorm: "primaryKey" json:"id"`
// CreateAt time.Time `json:"created_at",omitempty`//CreatedAt time.Time `json:"created_at",omitempty`
// UpdateAt time.Time `json:"created_at",omitempty`

	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at",omitempty`
	UpdatedAt time.Time `json:"updated_at",omitempty`

}