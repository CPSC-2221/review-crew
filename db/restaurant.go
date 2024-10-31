package db

import (
	"github.com/gin-gonic/gin"
)

type Restaurant struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description"`
	IsPremium   bool   `json:"isPremium"`
}

func CreateRestaurant(restaurant *Restaurant, c *gin.Context) (*Restaurant, error) {
	var new_restaurant Restaurant
	row := dbpool.QueryRow(c, "INSERT INTO restaurants(restaurantID, name, location, description, isPremium) VALUES (DEFAULT, $2, $3, $4, $5) RETURNING *;", restaurant.ID, restaurant.Name, restaurant.Location, restaurant.Description, restaurant.IsPremium)
	err := row.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Location, &restaurant.Description, &restaurant.IsPremium)
	if err != nil {
		return nil, err
	}
	return &new_restaurant, nil
}

func GetRestaurant(id int32, c *gin.Context) (*Restaurant, error) {
	var new_restaurant Restaurant
	row := dbpool.QueryRow(c, "SELECT * FROM restaurants WHERE restaurantID = $1;", id)
	err := row.Scan(&new_restaurant.ID, &new_restaurant.Name, &new_restaurant.Location, &new_restaurant.Description, &new_restaurant.IsPremium)
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
		rows.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Location, &restaurant.Description, &restaurant.IsPremium)
		restaurants = append(restaurants, restaurant)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func DeleteRestaurant(id int32, c *gin.Context) (*Restaurant, error) {
	var deleted_restaurant Restaurant
	row := dbpool.QueryRow(c, "DELETE FROM restaurants WHERE restaurantID = $1 RETURNING *;", id)
	err := row.Scan(&deleted_restaurant.ID, &deleted_restaurant.Name, &deleted_restaurant.Location, &deleted_restaurant.Description, &deleted_restaurant.IsPremium)
	if err != nil {
		return nil, err
	}
	return &deleted_restaurant, nil
}

func UpdateRestaurant(replaceWith *Restaurant, id int32, c *gin.Context) (*Restaurant, error) {
	var new_restaurant Restaurant
	row := dbpool.QueryRow(c, "UPDATE restaurants SET name=$2, location=$3, description=$4, isPremium=$5 where restaurantID=$1 RETURNING *;", id, replaceWith.Name, replaceWith.Location, replaceWith.Description, replaceWith.IsPremium)
	err := row.Scan(&new_restaurant.ID, &new_restaurant.Name, &new_restaurant.Location, &new_restaurant.Description, &new_restaurant.IsPremium)
	if err != nil {
		return nil, err
	}
	return &new_restaurant, nil
}
