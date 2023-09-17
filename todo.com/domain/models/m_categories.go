package models

import (
	"time"
)

// m_categories tables
type MCategories struct {
	ID        uint32    `gorm:"type:int;primarykey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	ColorCode string    `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time `gorm:"type:datetime(6);not null"`
	UpdatedAt time.Time `gorm:"type:datetime(6)"`
}
