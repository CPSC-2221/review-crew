package db

import (
	"github.com/gin-gonic/gin"
)

type Manages struct {
	Email             string `json:"email"`
	RestaurantID      int    `json:"restaurantID"`
	CanDeleteComments bool   `json:"candeletecomments"`
	CanUpdateListing  bool   `json:"canupdatelisting"`
}

func CreateManages(manages *Manages, c *gin.Context) (*Manages, error) {
	var new_manages Manages
	row := dbpool.QueryRow(c, "INSERT INTO manages(email,restaurantID,CanDeleteComments,CanUpdateListing) VALUES ($1,$2,$3,$4) RETURNING *;", manages.Email, manages.RestaurantID, manages.CanDeleteComments, manages.CanUpdateListing)
	err := row.Scan(&new_manager.Email, &new_manager.RestaurantID, &new_manager.CanDeleteComments, &new_manager.CanUpdateListing)
	if err != nil {
		return nil, err
	}
	return &new_manages, nil
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

func DeleteManages(email string, restaurantID int32, c *gin.Context) (*Manages, error) {
	var deleted_manages Manages
	row := dbpool.QueryRow(c, "DELETE FROM manages WHERE email = $1 AND restaurantID = $2 RETURNING *;", email, restaurantID)
	err := row.Scan(&deleted_manages.Email, &deleted_manages.RestaurantID, &deleted_manages.CanDeleteComments, &deleted_manages.CanUpdateListing)

	if err != nil {
		return nil, err
	}
	return &deleted_manages, nil
}

func UpdateManages(manages *Manages, c *gin.Context) (*Manages, error){
	var new_manages Manages
	row := dbpool.QueryRow(c, "UPDATE manages SET email=$1, restaurantID=$2, CanDeleteComments=$3, CanUpdateListing=$4 where email=$1 AND restaurantID=$2 RETURNING *;",manages.Email,manages.RestaurantID,manages.CanDeleteComments,manages.CanUpdateListing)
	err := row.Scan(&new_manages.Email,&new_manages.RestaurantID,&new_manages.CanDeleteComments,&new_manages.CanUpdateListing)
	if err != nil {
		return nil, err
	}
	return &new_manages, nil

}
