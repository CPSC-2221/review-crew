package db

import (
	"github.com/gin-gonic/gin"
)

type ReviewCharacterLimit struct {
	RestaurantID   int `json:"restaurantID"`
	CharacterLimit int `json:"characterLimit"`
}

func CreateReviewCharacterLimit(limit *ReviewCharacterLimit, c *gin.Context) (*ReviewCharacterLimit, error) {
	var new_limit ReviewCharacterLimit
	row := dbpool.QueryRow(c, "INSERT INTO review_character_limits(restaurantID, characterLimit) VALUES ($1, $2) RETURNING *;",
		limit.RestaurantID, limit.CharacterLimit)
	err := row.Scan(&new_limit.RestaurantID, &new_limit.CharacterLimit)
	if err != nil {
		return nil, err
	}
	return &new_limit, nil
}

func GetReviewCharacterLimit(restaurantID int32, c *gin.Context) (*ReviewCharacterLimit, error) {
	var limit ReviewCharacterLimit
	row := dbpool.QueryRow(c, "SELECT * FROM review_character_limits WHERE restaurantID = $1;", restaurantID)
	err := row.Scan(&limit.RestaurantID, &limit.CharacterLimit)
	if err != nil {
		return nil, err
	}
	return &limit, nil
}

func DeleteReviewCharacterLimit(restaurantID int, c *gin.Context) (*ReviewCharacterLimit, error) {
	var deleted_limit ReviewCharacterLimit
	row := dbpool.QueryRow(c, "DELETE FROM review_character_limits WHERE restaurantID = $1 RETURNING *;", restaurantID)
	err := row.Scan(&deleted_limit.RestaurantID, &deleted_limit.CharacterLimit)
	if err != nil {
		return nil, err
	}
	return &deleted_limit, nil
}

func UpdateReviewCharacterLimit(limit *ReviewCharacterLimit, c *gin.Context) (*ReviewCharacterLimit, error) {
	var updated_limit ReviewCharacterLimit
	row := dbpool.QueryRow(c, "UPDATE review_character_limits SET characterLimit=$1 WHERE restaurantID=$2 RETURNING *;",
		limit.CharacterLimit, limit.RestaurantID)
	err := row.Scan(&updated_limit.RestaurantID, &updated_limit.CharacterLimit)
	if err != nil {
		return nil, err
	}
	return &updated_limit, nil
}
