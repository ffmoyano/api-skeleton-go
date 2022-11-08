package dao

import (
	"database/sql"
	"notas/src/database"
	"notas/src/entity"
)

var err error

func GetUsers() ([]entity.User, error) {
	db := database.Get()
	var users []entity.User
	var result *sql.Rows
	var innerResult *sql.Rows

	// queries to the database
	result, err = db.Query(
		"Select id, name, email, created_at from user where verified = 1")
	if err != nil {
		return users, err
	}

	for result.Next() {
		var user entity.User
		var roles []entity.Role
		err = result.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			return users, err
		}

		innerResult, err = db.Query(
			"select r.id, r.name from role r inner join user_role ur on r.id = ur.role_id where ur.user_id = ?",
			user.Id)
		if err != nil {
			return users, err
		}

		for innerResult.Next() {
			var role entity.Role
			err = innerResult.Scan(&role.Id, &role.Name)
			if err != nil {
				return users, err
			}
			roles = append(roles, role)
		}
		user.Roles = roles
		users = append(users, user)
	}

	if innerResult != nil {
		if err = innerResult.Close(); err != nil {
			return users, err
		}
	}

	if result != nil {
		if err = result.Close(); err != nil {
			return users, err
		}
	}

	return users, nil
}
