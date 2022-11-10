package entity

import (
	"notas/src/model/dto"
)

type Role struct {
	Id   int
	Name string
}

func (role Role) ToDto() dto.Role {

	return dto.Role{
		Id:   role.Id,
		Name: role.Name,
	}

}
