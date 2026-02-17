package models

import "time"

type Game struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Code          string    `json:"code" gorm:"uniqueIndex; not null; type:varchar(255)"`
	CurrencyID    uint      `json:"currency_id" gorm:"index; not null"`
	CreatorID     uint      `json:"creator_id" gorm:"index; not null"`
	Bet           uint      `json:"bet" gorm:"not null"`
	WinningPoints uint      `json:"winning_points" gorm:"not null"`
	JoinType      string    `json:"join_type" gorm:"type:varchar(255); default:'anyone'; not null; index"`
	StartedAt     time.Time `json:"started_at"`
	FinishedAt    time.Time `json:"finished_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Currency      Currency  `json:"currency"`
	Creator       User      `json:"creator"`
	Users         []User    `json:"users" gorm:"many2many:game_user;"`
}

type GameUser struct {
	UserID   uint `json:"user_id" gorm:"primaryKey; index; not null"`
	GameID   uint `json:"game_id" gorm:"primaryKey; index; not null"`
	IsWinner bool `json:"is_winner" gorm:"index; not null; default:false; type:boolean"`
}

func (GameUser) TableName() string {
	return "game_user"
}
