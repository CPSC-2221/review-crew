package db

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func CreateUser(user User, c *gin.Context) (*User, error) {
	var new_user User
	row := dbpool.QueryRow(c, "INSERT INTO users(username, email) VALUES ($1, $2) RETURNING *;", user.Username, user.Email)
	err := row.Scan(&new_user.Username, &new_user.Email)
	if err != nil {
		return nil, err
	}
	return &new_user, nil
}

func GetUser(email string, c *gin.Context) (*User, error) {
	var user User
	row := dbpool.QueryRow(c, "SELECT * FROM users WHERE email = $1;", email)
	err := row.Scan(&user.Email, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers(c *gin.Context) ([]User, error) {
	rows, err := dbpool.Query(c, "SELECT * FROM users;")
	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.Email, &user.Username)
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func DeleteUser(email string, c *gin.Context) (*User, error) {
	var deleted_user User
	row := dbpool.QueryRow(c, "DELETE FROM users WHERE email = $1 RETURNING *;", email)
	err := row.Scan(&deleted_user.Username, &deleted_user.Email)
	if err != nil {
		return nil, err
	}
	return &deleted_user, nil
}

func UpdateUser(replaceWith User, oldEmail string, c *gin.Context) (*User, error) {
	var new_user User
	row := dbpool.QueryRow(c, "UPDATE users SET username=$1, email=$2 where email=$3 RETURNING *;", replaceWith.Username, replaceWith.Email, oldEmail)
	err := row.Scan(&new_user.Username, &new_user.Email)
	if err != nil {
		return nil, err
	}
	return &new_user, nil
}
