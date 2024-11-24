package db

import (
	"github.com/gin-gonic/gin"
)

type Own struct {
	Email        string `json:"email"`
	RestaurantID int32  `json:"restaurantID"`
}

func CreateOwn(owns Own, c *gin.Context) (*Own, error) {
	var new_own Own
	row := dbpool.QueryRow(c, "INSERT INTO owns(email, restaurantID) VALUES ($1, $2) RETURNING *;", owns.Email, owns.RestaurantID)
	err := row.Scan(&new_own.Email, &new_own.RestaurantID)
	if err != nil {
		return nil, err
	}
	return &new_own, nil
}

func GetUsersThatCanEditDesciption(restaurantID int32, c *gin.Context) ([]string, error) {
	rows, err := dbpool.Query(c, "SELECT email FROM owns WHERE restaurantID = $1 UNION SELECT email from manages WHERE restaurantID = $1 AND canUpdateListing;", restaurantID)
	if err != nil {
		return nil, err
	}

	var emails []string
	for rows.Next() {
		var email string
		rows.Scan(&email)
		emails = append(emails, email)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return emails, nil
}

func GetUsersThatCanDeleteReviews(restaurantID int32, c *gin.Context) []string {
	rows, err := dbpool.Query(c, "SELECT email FROM owns WHERE restaurantID = $1 UNION SELECT email from manages WHERE restaurantID = $1 AND canDeleteComments;", restaurantID)
	if err != nil {
		return nil
	}

	var emails []string
	for rows.Next() {
		var email string
		rows.Scan(&email)
		emails = append(emails, email)
	}
	err = rows.Err()
	if err != nil {
		return nil
	}
	return emails
}

func IsUserOwner(email string, c *gin.Context) bool {
	row := dbpool.QueryRow(c, "SELECT COUNT(*) FROM owns WHERE email = $1;", email)
	var ret int
	err := row.Scan(&ret)
	if err != nil {
		return false
	}
	return ret > 0
}
