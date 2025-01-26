package models

import (
	"errors"
	"rest_api/data_access"
	"rest_api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := data_access.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = id
	return nil
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := data_access.DB.QueryRow(query, u.Email)

	storedPassword := ""
	err := row.Scan(&u.ID, &storedPassword)
	if err != nil {
		return err
	}

	isPasswordValid := utils.ValidatePassword(u.Password, storedPassword)
	if !isPasswordValid {
		return errors.New("invalid password")
	}

	return nil
}
