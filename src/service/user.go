package service

import (
	"notas/src/dao"
	"notas/src/dto"
	"notas/src/entity"
)

func GetUsers() ([]dto.User, error) {
	users, err := dao.GetUsers()
	if err != nil {
		return []dto.User{}, err
	}
	return mapUsers(users), err

}

func mapUsers(users []entity.User) []dto.User {
	var usersDto []dto.User

	for _, user := range users {
		usersDto = append(usersDto, user.ToDto())
	}

	return usersDto
}
