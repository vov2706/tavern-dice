package responses

import "app/models"

type UserResponse struct {
	Data UserResource `json:"data"`
}

type UserResource struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func NewUserResource(user models.User) UserResource {
	return UserResource{
		ID:       user.ID,
		Username: user.Username,
	}
}
