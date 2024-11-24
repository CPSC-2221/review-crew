package router

import (
	"net/http"
	"server-api/db"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func createReview(ctx *gin.Context) {
	var review db.Review
	review.Email = ctx.PostForm("email")
	review.Comment = ctx.PostForm("comment")

	review.Datetime = time.Now()

	ridstr := ctx.PostForm("restaurantID")
	rid, _ := strconv.ParseInt(string(ridstr), 10, 32)
	review.RestaurantID = int(rid)

	_, err := db.CreateReview(review, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("HX-Refresh", "true")
}

func createReply(ctx *gin.Context) {
	var review db.Review
	review.Email = ctx.PostForm("email")
	review.Comment = ctx.PostForm("comment")

	review.Datetime = time.Now()

	ridstr := ctx.PostForm("restaurantID")
	reviewidstr := ctx.PostForm("reviewID")
	rid, _ := strconv.ParseInt(string(ridstr), 10, 32)
	reviewid, _ := strconv.ParseInt(string(reviewidstr), 10, 32)
	review.RestaurantID = int(rid)

	reviewResponse, err := db.CreateReview(review, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = db.CreateRepliesTo(ctx, db.RepliesTo{
		RepliesToReviewID:   reviewResponse.ID,
		IsRepliedToReviewID: int32(reviewid),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("HX-Refresh", "true")
}

func removeReview(ctx *gin.Context) {
	id_32, err := strconv.ParseInt(ctx.PostForm("reviewID"), 10, 32)
	if err != nil {
		panic(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := int32(id_32)

	var replies []db.NamedReview
	replies, _ = db.GetRepliesToAReview(ctx, id)
	for _, r := range replies {
		_, err = db.DeleteReview(r.ID, ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	_, err = db.DeleteReview(id, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}
