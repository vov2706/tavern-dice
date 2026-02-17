package responses

import "app/models"

type GameResource struct {
	ID            uint             `json:"id"`
	Code          string           `json:"code"`
	Bet           uint             `json:"bet"`
	WinningPoints uint             `json:"winning_points"`
	Link          string           `json:"link"`
	Currency      CurrencyResource `json:"currency"`
}

func NewGameResource(game models.Game) GameResource {
	return GameResource{
		ID:            game.ID,
		Code:          game.Code,
		Currency:      NewCurrencyResource(game.Currency),
		Bet:           game.Bet,
		WinningPoints: game.WinningPoints,
		Link:          game.Code,
	}
}
