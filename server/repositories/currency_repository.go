package repositories

import (
	"app/database"
	"app/models"
	"context"

	"gorm.io/gorm"
)

type CurrencyRepository struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *CurrencyRepository {
	return &CurrencyRepository{db: db}
}

func (repo *CurrencyRepository) FindById(currencyId uint) (models.Currency, error) {
	ctx := context.Background()

	return gorm.G[models.Currency](database.DB).
		Where("id = ?", currencyId).
		First(ctx)
}
