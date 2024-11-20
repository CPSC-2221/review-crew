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

func postReview(ctx *gin.Context) {
	var review db.Review

	err := ctx.Bind(&review)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := db.CreateReview(review, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"review": res,
	})
}

func getReview(ctx *gin.Context) {
	id_32, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := int32(id_32)

	res, err := db.GetReview(id, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"review": res,
	})
}

func deleteReview(ctx *gin.Context) {
	id_32, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := int32(id_32)

	res, err := db.DeleteReview(id, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"deleted_review": res,
	})
}

func putReview(ctx *gin.Context) {
	var updated_review db.Review

	err := ctx.Bind(&updated_review)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id_32, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := int32(id_32)
	updated_review.ID = id

	res, err := db.UpdateReview(updated_review, id, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"updated_review": res,
	})
}
