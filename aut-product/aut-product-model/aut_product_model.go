package autproductmodel

import (
	"database/sql"
	"product/config"
	"product/entities"
	"product/helper"
)

func Login(user *entities.Users, username, password string) error {
	con := config.CreateCon()

	script := "SELECT * from users where username = ?"

	err := con.QueryRow(script, username).Scan(
		&user.Id, &user.Name, &user.Email, &user.Username, &user.Password,
	)
	if err != nil {
		return err
	}

	if err == sql.ErrNoRows {
		return err
	}

	match, err := helper.CheckPasswordHash(password, user.Password)
	if !match {
		return err
	}

	return nil
}
