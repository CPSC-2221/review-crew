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

type ReviewLikes struct {
	ReviewID int32
	Likes    int
}

func GetReviewLikesCountLocation(c *gin.Context, restaurantID int32) ([]ReviewLikes, error) {
	rows, err := dbpool.Query(c, "SELECT reviewID, Count(*) FROM likes WHERE reviewID IN (SELECT reviewID FROM review WHERE restaurantID = $1) GROUP BY reviewID;", restaurantID)
	if err != nil {
		return nil, err
	}

	var likes []ReviewLikes
	for rows.Next() {
		var review ReviewLikes
		err := rows.Scan(&review.ReviewID, &review.Likes)
		if err != nil {
			return nil, err
		}
		likes = append(likes, review)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return likes, nil
}

func GetUserLikedReviewIDsByLocation(c *gin.Context, email string, restaurantID int32) ([]int32, error) {
	rows, err := dbpool.Query(c, "SELECT reviewID FROM likes WHERE reviewID IN (SELECT reviewID FROM review WHERE restaurantID = $1) AND email = $2;", restaurantID, email)
	if err != nil {
		return nil, err
	}

	var reviewIDs []int32
	for rows.Next() {
		var reviewid int32
		err := rows.Scan(&reviewid)
		if err != nil {
			return nil, err
		}
		reviewIDs = append(reviewIDs, reviewid)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return reviewIDs, nil
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
