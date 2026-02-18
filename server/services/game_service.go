package services

import (
	"app/http/inputs"
	"app/models"
	"app/repositories"
	"errors"
)

type GameService struct {
	balanceRepo  *repositories.BalanceRepository
	currencyRepo *repositories.CurrencyRepository
	gameRepo     *repositories.GameRepository
}

func NewGameService(
	balanceRepo *repositories.BalanceRepository,
	currencyRepo *repositories.CurrencyRepository,
	gameRepo *repositories.GameRepository,
) *GameService {
	return &GameService{
		balanceRepo:  balanceRepo,
		currencyRepo: currencyRepo,
		gameRepo:     gameRepo,
	}
}

func (service *GameService) CreateGame(authUser *models.User, input *inputs.CreateGameInput) (*models.Game, error) {
	currency, err := service.currencyRepo.FindById(input.CurrencyID)

	if err != nil {
		return nil, errors.New("failed to get currency")
	}

	userBalance, err := service.balanceRepo.FindByUserAndCurrency(*authUser, currency.ID)

	if err != nil {
		return nil, errors.New("failed to get user balance")
	}

	if userBalance.Amount < input.Bet {
		return nil, errors.New("insufficient funds")
	}

	game, err := service.gameRepo.CreateGame(*authUser, *input)

	if err != nil {
		return nil, errors.New("failed to create game")

	}

	game.Currency = currency

	return game, nil
}
