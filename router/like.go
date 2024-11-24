package router

import (
	"net/http"
	"server-api/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
