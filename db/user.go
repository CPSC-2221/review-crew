package db

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type KeylessUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type User struct {
	ID int16 `json:"id"`
	KeylessUser
}

func CreateUsersTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE users ("+
			"id SMALLSERIAL PRIMARY KEY,"+
			"username VARCHAR (50) UNIQUE NOT NULL,"+
			"email VARCHAR (255) UNIQUE NOT NULL"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateUser(user *KeylessUser, c *gin.Context) (*User, error) {
	var new_user User
	row := dbpool.QueryRow(c, "INSERT INTO users(id, username, email) VALUES (DEFAULT, $1, $2) RETURNING *;", user.Username, user.Email)
	err := row.Scan(&new_user.ID, &new_user.Username, &new_user.Email)
	if err != nil {
		return nil, err
	}
	return &new_user, nil
}

func GetUser(id int16, c *gin.Context) (*User, error) {
	var user User
	row := dbpool.QueryRow(c, "SELECT * FROM users WHERE id = $1;", id)
	err := row.Scan(&user.ID, &user.Username, &user.Email)
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
		rows.Scan(&user.ID, &user.Username, &user.Email)
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func DeleteUser(id int16, c *gin.Context) (*User, error) {
	var deleted_user User
	row := dbpool.QueryRow(c, "DELETE FROM users WHERE id = $1 RETURNING *;", id)
	err := row.Scan(&deleted_user.ID, &deleted_user.Username, &deleted_user.Email)
	if err != nil {
		return nil, err
	}
	return &deleted_user, nil
}

func UpdateUser(user *User, c *gin.Context) (*User, error) {
	var new_user User
	row := dbpool.QueryRow(c, "UPDATE users SET username=$1, email=$2 where id = $2 RETURNING *;", user.Username, user.Email, user.ID)
	err := row.Scan(&new_user.ID, &new_user.Username, &new_user.Email)
	if err != nil {
		return nil, err
	}
	return &new_user, nil
}
