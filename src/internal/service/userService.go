package service

import (
	"notas/src/database"
	"notas/src/model/dto"
	"notas/src/model/entity"
)

func GetUsers() ([]dto.User, error) {
	users, err := database.GetUsers()
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
