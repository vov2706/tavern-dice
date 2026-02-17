package models

import "time"

const (
	GOLD   = "gold"
	SILVER = "silver"
	BRONZE = "bronze"
)

type Currency struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Slug      string    `json:"slug" gorm:"uniqueIndex; not null; type:varchar(255)"`
	Name      string    `json:"name" gorm:"varchar(255); not null; type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
