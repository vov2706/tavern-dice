package repositories

import (
	"app/database"
	"app/http/inputs"
	"app/models"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) *GameRepository {
	return &GameRepository{db: db}
}

func (repo *GameRepository) CreateGame(user models.User, input inputs.CreateGameInput) (*models.Game, error) {
	ctx := context.Background()
	game := models.Game{
		Code:          uuid.New().String(),
		CreatorID:     user.ID,
		CurrencyID:    input.CurrencyID,
		Bet:           input.Bet,
		WinningPoints: input.WinningPoints,
		JoinType:      input.JoinType,
	}

	err := gorm.G[models.Game](database.DB).Create(ctx, &game)

	return &game, err
}
