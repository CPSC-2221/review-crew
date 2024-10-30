package router

import (
	"net/http"
	"server-api/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func postReview(ctx *gin.Context) {
	var review db.Review

	err := ctx.Bind(&review)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error1": err.Error(),
		})
		return
	}

	res, err := db.CreateReview(&review, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error2": err.Error(),
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

	dbReview, err := db.GetReview(id, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dbReview.ID = updated_review.ID
	dbReview.Email = updated_review.Email
	dbReview.Comment = updated_review.Comment
	dbReview.Date = updated_review.Date

	res, err := db.UpdateReview(dbReview, ctx)
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
