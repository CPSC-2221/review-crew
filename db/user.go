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
	err := row.Scan(&new_user.Email, &new_user.Username)
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

func GetEmailFromUsername(username string, c *gin.Context) string {
	var email string
	row := dbpool.QueryRow(c, "SELECT email FROM users WHERE username = $1;", username)
	row.Scan(&email)
	return email
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

func GetUsersLikedEveryReview(restaurantID int32, c *gin.Context) []User {
	rows, err := dbpool.Query(c, "SELECT * FROM users u WHERE NOT EXISTS (SELECT r.reviewID FROM review r WHERE r.restaurantID = $1 AND r.reviewID NOT IN (SELECT repliesToReviewID FROM repliesTo) AND NOT EXISTS (SELECT l.reviewID, l.email FROM likes l WHERE l.reviewID = r.reviewID AND l.email = u.email));", restaurantID)
	if err != nil {
		println("query")
		panic(err)
		return nil
	}

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.Email, &user.Username)
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		println("assignment")
		panic(err)
		return nil
	}
	return users
}

func DeleteUser(email string, c *gin.Context) (*User, error) {
	var deleted_user User
	row := dbpool.QueryRow(c, "DELETE FROM users WHERE email = $1 RETURNING *;", email)
	err := row.Scan(&deleted_user.Email, &deleted_user.Username)
	if err != nil {
		return nil, err
	}
	return &deleted_user, nil
}

func UpdateUser(replaceWith User, oldEmail string, c *gin.Context) (*User, error) {
	var new_user User
	row := dbpool.QueryRow(c, "UPDATE users SET username=$1, email=$2 where email=$3 RETURNING *;", replaceWith.Username, replaceWith.Email, oldEmail)
	err := row.Scan(&new_user.Email, &new_user.Username)
	if err != nil {
		return nil, err
	}
	return &new_user, nil
}
