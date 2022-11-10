package entity

import (
	dto2 "notas/src/model/dto"
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

func (user User) ToDto() dto2.User {
	var rolesDto []dto2.Role
	for _, role := range user.Roles {
		rolesDto = append(rolesDto, role.ToDto())
	}
	return dto2.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Roles:     rolesDto,
		CreatedAt: user.CreatedAt,
	}

}
