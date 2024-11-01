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

func GetOwn(email string, restaurantID int32, c *gin.Context) (*Own, error) {
	var own Own
	row := dbpool.QueryRow(c, "SELECT * FROM owns WHERE email = $1 and restaurantID = $2;", email, restaurantID)
	err := row.Scan(&own.Email, &own.RestaurantID)
	if err != nil {
		return nil, err
	}
	return &own, nil
}

func GetOwns(c *gin.Context) ([]Own, error) {
	rows, err := dbpool.Query(c, "SELECT * FROM owns;")
	if err != nil {
		return nil, err
	}

	var owns []Own
	for rows.Next() {
		var own Own
		rows.Scan(&own.Email, &own.RestaurantID)
		owns = append(owns, own)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return owns, nil
}

func DeleteOwn(email string, restaurantID int32, c *gin.Context) (*Own, error) {
	var deleted_own Own
	row := dbpool.QueryRow(c, "DELETE FROM owns WHERE email = $1 and restaurantID = $2 RETURNING *;", email, restaurantID)
	err := row.Scan(&deleted_own.Email, &deleted_own.RestaurantID)
	if err != nil {
		return nil, err
	}
	return &deleted_own, nil
}

func UpdateOwn(replaceWith Own, oldEmail string, oldRestaurantID int32, c *gin.Context) (*Own, error) {
	var new_own Own
	row := dbpool.QueryRow(c, "UPDATE owns SET email=$1, restaurantID=$2 where email = $3 and restaurantID = $4 RETURNING *;", replaceWith.Email, replaceWith.RestaurantID, oldEmail, oldRestaurantID)
	err := row.Scan(&new_own.Email, &new_own.RestaurantID)
	if err != nil {
		return nil, err
	}
	return &new_own, nil
}
