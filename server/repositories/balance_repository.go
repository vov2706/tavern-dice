package repositories

import (
	"app/models"
	"context"

	"gorm.io/gorm"
)

type BalanceRepository struct {
	db *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) *BalanceRepository {
	return &BalanceRepository{db: db}
}

func (repo *BalanceRepository) FindByUserAndCurrency(user models.User, currencyId uint) (*models.Balance, error) {
	if len(user.Balances) > 0 {
		for _, b := range user.Balances {
			if b.CurrencyID == currencyId {
				return &b, nil
			}
		}
	}

	ctx := context.Background()
	balance, err := gorm.G[models.Balance](repo.db).
		Where("user_id = ?", user.ID).
		Where("currency_id = ?", currencyId).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return &balance, nil
}
