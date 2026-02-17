package responses

import "app/models"

type CurrencyResource struct {
	ID   uint   `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

func NewCurrencyResource(currency models.Currency) CurrencyResource {
	return CurrencyResource{
		ID:   currency.ID,
		Slug: currency.Slug,
		Name: currency.Name,
	}
}
