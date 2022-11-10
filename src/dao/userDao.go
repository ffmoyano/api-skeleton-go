package dao

import (
	"database/sql"
	"notas/src/database"
	"notas/src/entity"
	"notas/src/logger"
)

var err error

func GetUsers() ([]entity.User, error) {
	db := database.Get()
	var users []entity.User
	var rows *sql.Rows

	rows, err = db.Query(
		"Select id, name, email, created_at from user where verified = 1")
	if err != nil {
		logger.Error(err.Error())
		return users, err
	}

	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			logger.Error(err.Error())
			return users, err
		}

		user.Roles, err = GetRolesByUser(user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	err = rows.Close()
	if err != nil {
		logger.Warn(err.Error())
	}

	return users, nil
}
