package models

import (
	"EventBooking/db"
	"EventBooking/utils"
	"errors"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	haspwd, err := utils.HashPassword(u.Password)

	result, err := stmt.Exec(u.Email, haspwd)
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

func (u User) Login() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var id int64
	var password string
	err := row.Scan(&id, &password)
	if err != nil {
		return err
	}
	return nil

}

func (u User) ValidateUser() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}
	validPassword := utils.CheckPassword(u.Password, retrievedPassword)
	if !validPassword {
		return errors.New("Invalid Credentials")
	}

	return nil
}
