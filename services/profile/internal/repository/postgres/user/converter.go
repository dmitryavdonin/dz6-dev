package repository

import (
	"profile/internal/domain/user"
	"profile/internal/repository/postgres/user/dao"
)

func (r *Repository) toDomainUser(daoUser *dao.User) (*user.User, error) {
	return user.NewUserWithId(daoUser.Id, daoUser.Login, daoUser.Password, daoUser.Name, daoUser.Middlename, daoUser.Surname, daoUser.Phone, daoUser.City, daoUser.Role, daoUser.CreatedAt, daoUser.ModifiedAt), nil
}
