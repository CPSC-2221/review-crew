package db

import (
	"github.com/gin-gonic/gin"
)

type Manages struct {
	RestaurantID      int32  `json:"restaurantID"`
	Email             string `json:"email"`
	CanDeleteComments bool   `json:"canDeleteComments"`
	CanUpdateListing  bool   `json:"canUpdateListing"`
}

func CreateManages(manages Manages, c *gin.Context) {
	dbpool.QueryRow(c, "INSERT INTO manages(email,restaurantID,canDeleteComments,canUpdateListing) VALUES ($1,$2,$3,$4);", manages.Email, manages.RestaurantID, manages.CanDeleteComments, manages.CanUpdateListing)
}

func GetManages(email string, restaurantID int32, c *gin.Context) (*Manages, error) {
	var manages Manages
	row := dbpool.QueryRow(c, "SELECT * FROM manages WHERE email = $1 AND restaurantID = $2;", email, restaurantID)
	err := row.Scan(&manages.Email, &manages.RestaurantID, &manages.CanDeleteComments, &manages.CanUpdateListing)
	if err != nil {
		return nil, err
	}
	return &manages, nil
}

type Manager struct {
	Username          string
	Email             string
	CanDeleteComments bool
	CanUpdateListing  bool
}

func GetRestaurantManagers(restaurantID int32, c *gin.Context) []Manager {
	rows, err := dbpool.Query(c, "SELECT u.Username, m.email, m.CanDeleteComments, m.CanUpdateListing FROM users u, manages m WHERE u.email = m.email AND m.restaurantID = $1;", restaurantID)
	if err != nil {
		panic(err)
		return nil
	}

	var managers []Manager
	for rows.Next() {
		var manager Manager
		err := rows.Scan(&manager.Username, &manager.Email, &manager.CanDeleteComments, &manager.CanUpdateListing)
		if err != nil {
			panic(err)
			return nil
		}
		managers = append(managers, manager)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
		return nil
	}
	return managers
}

func DeleteManages(email string, restaurantID int32, c *gin.Context) {
	dbpool.QueryRow(c, "DELETE FROM manages WHERE email = $1 AND restaurantID = $2;", email, restaurantID)
}

func UpdateManages(manages Manages, c *gin.Context) (*Manages, error) {
	var new_manages Manages
	row := dbpool.QueryRow(c, "UPDATE manages SET CanDeleteComments=$3, CanUpdateListing=$4 where email=$1 AND restaurantID=$2 RETURNING *;", manages.Email, manages.RestaurantID, manages.CanDeleteComments, manages.CanUpdateListing)
	err := row.Scan(&new_manages.Email, &new_manages.RestaurantID, &new_manages.CanDeleteComments, &new_manages.CanUpdateListing)
	if err != nil {
		return nil, err
	}
	return &new_manages, nil
}
