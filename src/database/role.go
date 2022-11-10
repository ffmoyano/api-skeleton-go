package database

import (
	"database/sql"
	"notas/src/internal/logger"
	entity2 "notas/src/model/entity"
)

func GetRolesByUser(user entity2.User) ([]entity2.Role, error) {
	var roles []entity2.Role
	var rows *sql.Rows

	rows, err = db.Query(
		"select r.id, r.name from role r inner join user_role ur on r.id = ur.role_id where ur.user_id = ?",
		user.Id)
	if err != nil {
		logger.Error(err.Error())
		return roles, err
	}
	for rows.Next() {
		var role entity2.Role
		err = rows.Scan(&role.Id, &role.Name)
		if err != nil {
			logger.Error(err.Error())
			return roles, err
		}
		roles = append(roles, role)
	}

	if rows != nil {
		err = rows.Close()
		if err != nil {
			logger.Warn(err.Error())
			return roles, err
		}
	}

	return roles, nil
}
