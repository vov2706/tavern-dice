package models

import "time"

type Balance struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id" gorm:"index; not null"`
	CurrencyID uint      `json:"currency_id" gorm:"index; not null"`
	Amount     uint      `json:"amount" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	User       User      `gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
	Currency   Currency  `gorm:"foreignKey:CurrencyID; constraint:OnDelete:CASCADE"`
}
