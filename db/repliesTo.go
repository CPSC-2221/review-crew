package db

import (
	"github.com/gin-gonic/gin"
)

type RepliesTo struct {
	RepliesToReviewID   int32 `json:"repliesTo"`
	IsRepliedToReviewID int32 `json:"isRepliedTo"`
}

func CreateRepliesTo(c *gin.Context, repliesTo RepliesTo) (*RepliesTo, error) {
	var new_repliesTo RepliesTo
	row := dbpool.QueryRow(
		c,
		"INSERT INTO repliesTo(repliesToReviewID,isRepliedToReviewID) VALUES ($1,$2) RETURNING *;",
		repliesTo.RepliesToReviewID,
		repliesTo.IsRepliedToReviewID,
	)
	err := row.Scan(&new_repliesTo.RepliesToReviewID, &new_repliesTo.IsRepliedToReviewID)
	if err != nil {
		return nil, err
	}
	return &new_repliesTo, nil
}

func CountRepliesToReview(c *gin.Context, reviewID int32) int {
	row := dbpool.QueryRow(c, "SELECT Count(*) FROM repliesTo WHERE isRepliedToReviewID = $1;", reviewID)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

func DeleteRepliesTo(c *gin.Context, replyToReviewID int32, isRepliedToReviewID int32) (*RepliesTo, error) {
	var deleted_repliesTo RepliesTo
	row := dbpool.QueryRow(c, "DELETE FROM repliesTo WHERE repliesToReviewID = $1 AND isRepliedToReviewID = $2 RETURNING *;", replyToReviewID, isRepliedToReviewID)
	err := row.Scan(&deleted_repliesTo.RepliesToReviewID, &deleted_repliesTo.IsRepliedToReviewID)
	if err != nil {
		return nil, err
	}
	return &deleted_repliesTo, nil
}