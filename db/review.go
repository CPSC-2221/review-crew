package db

import (
	"github.com/gin-gonic/gin"
)

type Review struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	RestaurantID int    `json:"restaurantID"`
	Comment      string `json:"comment"`
	Date         int    `json:"date"`
}

func CreateReview(review *Review, c *gin.Context) (*Review, error) {
	var new_review Review
	row := dbpool.QueryRow(c, "INSERT INTO review(id,email,restaurantID,comment,date) VALUES(DEFAULT,$1,$2,$3,$4);", review.ID, review.Email, review.RestaurantID, review.Comment, review.Date)
	err := row.Scan(&new_review.ID, &new_review.Email, &new_review.RestaurantID, &new_review.Comment, &new_review.Date)
	if err != nil {
		return nil, err
	}
	return &new_review, nil
}

func GetReview(id int32, c *gin.Context) (*Review, error) {
	var review Review
	row := dbpool.QueryRow(c, "SELECT * FROM review WHERE id = $1;", id)
	err := row.Scan(&review.ID, &review.Email, &review.RestaurantID, &review.Comment, &review.Date)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func GetReviews(c *gin.Context) ([]Review, error) {
	rows, err := dbpool.Query(c, "SELECT * FROM review;")
	if err != nil {
		return nil, err
	}
	var reviews []Review
	for rows.Next() {
		var review Review
		err := rows.Scan(&review.ID, &review.Email, &review.RestaurantID, &review.Comment, &review.Date)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func DeleteReview(id int32, c *gin.Context) (*Review, error) {
	var deleted_review Review
	row := dbpool.QueryRow(c, "DELETE FROM review WHERE id = $1 RETURNING *;", id)
	err := row.Scan(&deleted_review.ID, &deleted_review.Email, &deleted_review.RestaurantID, &deleted_review.Comment, &deleted_review.Date)
	if err != nil {
		return nil, err
	}
	return &deleted_review, nil
}

func UpdateReview(review *Review, c *gin.Context) (*Review, error) {
	var updated_review Review
	row := dbpool.QueryRow(c, "UPDATE review SET email=$1, restaurantID=$2, comment=$3, date=$4 WHERE id=$5 RETURNING *;",
		review.Email, review.RestaurantID, review.Comment, review.Date, review.ID)
	err := row.Scan(&updated_review.ID, &updated_review.Email, &updated_review.RestaurantID, &updated_review.Comment, &updated_review.Date)
	if err != nil {
		return nil, err
	}
	return &updated_review, nil
}
