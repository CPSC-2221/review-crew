package db

import (
	"github.com/gin-gonic/gin"
)

type Like struct {
	ReviewID int32  `json:"reviewID"`
	Email    string `json:"email"`
}

func CreateLike(like Like, c *gin.Context) (*Like, error) {
	var new_like Like
	row := dbpool.QueryRow(
		c,
		"INSERT INTO likes(reviewID,email) VALUES ($1,$2) RETURNING *;",
		like.ReviewID,
		like.Email,
	)
	err := row.Scan(&new_like.ReviewID, &new_like.Email)
	if err != nil {
		return nil, err
	}
	return &new_like, nil
}

func GetReviewLikes(reviewID int32, c *gin.Context) (*Like, error) {
	var like Like
	row := dbpool.QueryRow(c, "SELECT * FROM likes WHERE reviewID = $1;", reviewID)
	err := row.Scan(&like.ReviewID, &like.Email)
	if err != nil {
		return nil, err
	}
	return &like, nil
}

func DeleteLike(reviewID int32, email string, c *gin.Context) (*Like, error) {
	var like Like
	row := dbpool.QueryRow(
		c,
		"DELETE FROM likes WHERE reviewID = $1 and email = $2 RETURNING *;",
		reviewID,
		email,
	)
	err := row.Scan(&like.ReviewID, &like.Email)
	if err != nil {
		return nil, err
	}
	return &like, nil
}
