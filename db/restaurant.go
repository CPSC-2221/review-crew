package db

import (
	"github.com/gin-gonic/gin"
)

type Restaurant struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

func CreateRestaurant(restaurant *Restaurant, c *gin.Context) (*Restaurant, error) {
	var new_restaurant Restaurant
	row := dbpool.QueryRow(c, "INSERT INTO restaurants(restaurantID, name, location, description) VALUES (DEFAULT, $1, $2, $3) RETURNING *;", restaurant.Name, restaurant.Location, restaurant.Description)
	err := row.Scan(&new_restaurant.ID, &new_restaurant.Name, &new_restaurant.Location, &new_restaurant.Description)
	if err != nil {
		return nil, err
	}
	return &new_restaurant, nil
}

func GetRestaurant(id int32, c *gin.Context) (*Restaurant, error) {
	var new_restaurant Restaurant
	row := dbpool.QueryRow(c, "SELECT * FROM restaurants WHERE restaurantID = $1;", id)
	err := row.Scan(&new_restaurant.ID, &new_restaurant.Name, &new_restaurant.Location, &new_restaurant.Description)
	if err != nil {
		return nil, err
	}
	return &new_restaurant, nil
}

func GetRestaurants(c *gin.Context) ([]Restaurant, error) {
	rows, err := dbpool.Query(c, "SELECT * FROM restaurants;")
	if err != nil {
		return nil, err
	}

	var restaurants []Restaurant
	for rows.Next() {
		var restaurant Restaurant
		rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Location, &restaurant.Description)
		restaurants = append(restaurants, restaurant)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func GetUserOwnedRestaurants(email string, c *gin.Context) []Restaurant {
	rows, err := dbpool.Query(c, "SELECT * FROM restaurants WHERE restaurantID IN (SELECT restaurantID FROM owns WHERE email = $1);", email)
	if err != nil {
		return nil
	}

	var restaurants []Restaurant
	for rows.Next() {
		var restaurant Restaurant
		rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Location, &restaurant.Description)
		restaurants = append(restaurants, restaurant)
	}
	err = rows.Err()
	if err != nil {
		return nil
	}
	return restaurants
}

func UpdateDescription(replaceWith string, id int32, c *gin.Context) {
	dbpool.QueryRow(c, "UPDATE restaurants SET description=$1 where restaurantID=$2;", replaceWith, id)
}
