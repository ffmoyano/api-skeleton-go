package entity

import "notas/src/dto"

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
