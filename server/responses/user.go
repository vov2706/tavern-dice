package responses

import "app/models"

type UserResponse struct {
	Data UserResource `json:"data"`
}

type UserResource struct {
	ID       uint            `json:"id"`
	Username string          `json:"username"`
	Balance  BalanceResource `json:"balance"`
}

type BalanceResource struct {
	Amount   uint             `json:"amount"`
	Currency CurrencyResource `json:"currency"`
}

func NewUserResource(user models.User) UserResource {
	var balance BalanceResource

	if len(user.Balances) > 0 {
		for _, b := range user.Balances {
			balance = BalanceResource{
				Amount:   b.Amount,
				Currency: NewCurrencyResource(b.Currency),
			}

			break
		}
	}

	return UserResource{
		ID:       user.ID,
		Username: user.Username,
		Balance:  balance,
	}
}
