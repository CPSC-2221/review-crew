package db

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Review struct {
	ID           int32     `json:"id"`
	Email        string    `json:"email"`
	RestaurantID int       `json:"restaurantID"`
	Comment      string    `json:"comment"`
	Datetime     time.Time `json:"datetime"`
}

func CreateReview(review Review, c *gin.Context) (*Review, error) {
	var new_review Review
	row := dbpool.QueryRow(
		c,
		"INSERT INTO review(reviewID, email, restaurantID, comment, datetime) VALUES (DEFAULT, $1, $2, $3, $4) RETURNING *;",
		review.Email,
		review.RestaurantID,
		review.Comment,
		review.Datetime,
	)
	err := row.Scan(
		&new_review.ID,
		&new_review.Email,
		&new_review.RestaurantID,
		&new_review.Comment,
		&new_review.Datetime,
	)
	if err != nil {
		return nil, err
	}
	return &new_review, nil
}

func GetReview(id int32, c *gin.Context) (*Review, error) {
	var review Review
	row := dbpool.QueryRow(c, "SELECT * FROM review WHERE reviewID = $1;", id)
	err := row.Scan(&review.ID, &review.Email, &review.RestaurantID, &review.Comment, &review.Datetime)
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
		err := rows.Scan(&review.ID, &review.Email, &review.RestaurantID, &review.Comment, &review.Datetime)
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

type NamedReview struct {
	Review
	Username string
}

func GetRestaurantReviews(c *gin.Context, restaurantID int32) ([]NamedReview, error) {
	rows, err := dbpool.Query(c, "SELECT reviewID, r.email, restaurantID, comment, datetime, username FROM review r, users u where restaurantID = $1 AND reviewID NOT IN (SELECT repliesToReviewID FROM repliesTo) AND u.email = r.email;", restaurantID)
	if err != nil {
		return nil, err
	}
	var reviews []NamedReview
	for rows.Next() {
		var review NamedReview
		err := rows.Scan(&review.ID, &review.Email, &review.RestaurantID, &review.Comment, &review.Datetime, &review.Username)
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

func GetRepliesToAReview(c *gin.Context, reviewID int32) ([]NamedReview, error) {
	rows, err := dbpool.Query(c, "SELECT reviewID, r.email, restaurantID, comment, datetime, username FROM review r, users u WHERE reviewID IN (SELECT repliesToReviewID FROM repliesTo WHERE isRepliedToReviewID = $1) AND r.email = u.email;", reviewID)
	if err != nil {
		return nil, err
	}

	var reviews []NamedReview
	for rows.Next() {
		var review NamedReview
		err := rows.Scan(&review.ID, &review.Email, &review.RestaurantID, &review.Comment, &review.Datetime, &review.Username)
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

func CountReviewsOnRestaurant(c *gin.Context, restaurantID int32) int {
	row := dbpool.QueryRow(c, "SELECT Count(*) FROM review WHERE restaurantID = $1 AND reviewID NOT IN (SELECT repliesToReviewID FROM repliesTo);", restaurantID)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

func DeleteReview(id int32, c *gin.Context) (*Review, error) {
	var deleted_review Review
	row := dbpool.QueryRow(c, "DELETE FROM review WHERE reviewID = $1 RETURNING *;", id)
	err := row.Scan(&deleted_review.ID, &deleted_review.Email, &deleted_review.RestaurantID, &deleted_review.Comment, &deleted_review.Datetime)
	if err != nil {
		return nil, err
	}
	return &deleted_review, nil
}

func UpdateReview(review Review, id int32, c *gin.Context) (*Review, error) {
	var updated_review Review
	row := dbpool.QueryRow(c, "UPDATE review SET comment=$2 WHERE id=$1 RETURNING *;", id, review.Comment)
	err := row.Scan(&updated_review.ID, &updated_review.Email, &updated_review.RestaurantID, &updated_review.Comment, &updated_review.Datetime)
	if err != nil {
		return nil, err
	}
	return &updated_review, nil
}
