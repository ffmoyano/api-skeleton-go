package entity

import (
	"notas/src/dto"
	"time"
)

type User struct {
	Id        int
	Name      string
	Password  string
	Email     string
	Roles     []Role
	CreatedAt time.Time
	Verified  bool
}

func (user User) ToDto() dto.User {
	var rolesDto []dto.Role
	for _, role := range user.Roles {
		rolesDto = append(rolesDto, role.ToDto())
	}
	return dto.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Roles:     rolesDto,
		CreatedAt: user.CreatedAt,
	}

}
