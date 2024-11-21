package router

import (
	"net/http"
	"server-api/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func postLike(ctx *gin.Context) {
	var likes db.Like

	err := ctx.Bind(&likes)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := db.CreateLike(likes, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"like": res,
	})
}

func createLike(ctx *gin.Context) {
	reviewid_32, err := strconv.ParseInt(ctx.PostForm("reviewID"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	usr, _ := getUserFromCookie(ctx)

	_, err = db.CreateLike(db.Like{
		ReviewID: int32(reviewid_32),
		Email:    usr.Email,
	}, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("HX-Refresh", "true")
}

func removeLike(ctx *gin.Context) {
	reviewid_32, err := strconv.ParseInt(ctx.PostForm("reviewID"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	usr, _ := getUserFromCookie(ctx)

	_, err = db.DeleteLike(int32(reviewid_32), usr.Email, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("HX-Refresh", "true")
}

func getReviewLikes(ctx *gin.Context) {
	id_32, err := strconv.ParseInt(ctx.Param("reviewID"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	reviewID := int32(id_32)
	res, err := db.GetReviewLikes(reviewID, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"likes": res,
	})
}

func deleteLike(ctx *gin.Context) {
	email := ctx.Param("email")
	id_32, err := strconv.ParseInt(ctx.Param("reviewID"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	reviewID := int32(id_32)
	res, err := db.DeleteLike(reviewID, email, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"deleted_like": res,
	})
}
