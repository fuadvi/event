package convert

import (
	"rental_mobile_fiber/internal/entity"
	"rental_mobile_fiber/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {

	return &model.UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}
}
