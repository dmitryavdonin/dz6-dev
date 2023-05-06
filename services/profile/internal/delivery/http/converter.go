package delivery

import (
	response "profile/internal/delivery/http/user"
	"profile/internal/domain/user"
)

func (d *Delivery) toResponseUser(user *user.User) *response.UserResponse {
	return &response.UserResponse{
		Id:         user.Id(),
		Login:      user.Login(),
		Password:   "***********",
		Name:       user.Name(),
		Middlename: user.MiddleName(),
		Surname:    user.Surname(),
		Phone:      user.Phone(),
		City:       user.City(),
		Role:       user.Role(),
		CreatedAt:  user.CreatedAt(),
		ModifiedAt: user.ModifiedAt(),
	}
}
