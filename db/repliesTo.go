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

func GetRepliesToAReview(c *gin.Context, reviewID int32) ([]RepliesTo, error) {
	rows, err := dbpool.Query(c, "SELECT * FROM repliesTo WHERE isRepliedToReviewID = $1;", reviewID)
	if err != nil {
		return nil, err
	}

	var replies []RepliesTo
	for rows.Next() {
		var reply RepliesTo
		rows.Scan(&reply.RepliesToReviewID, &reply.IsRepliedToReviewID)
		replies = append(replies, reply)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return replies, nil
}

func IsReplyingToAReview(c *gin.Context, reviewID int32) bool {
	row := dbpool.QueryRow(c, "SELECT * FROM repliesTo WHERE repliesToReviewID = $1;", reviewID)
	if row == nil {
		return false
	}
	return true
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
