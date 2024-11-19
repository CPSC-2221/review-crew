package router

import (
	"net/http"
	"server-api/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func postRepliesTo(ctx *gin.Context) {
	var repliesTo db.RepliesTo

	err := ctx.Bind(&repliesTo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := db.CreateRepliesTo(ctx, repliesTo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"repliesTo": res,
	})
}

func deleteRepliesTo(ctx *gin.Context) {
	repliedTo_32, err := strconv.ParseInt(ctx.Param("repliedTo"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	repliesTo_32, err := strconv.ParseInt(ctx.Param("repliesTo"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	repliedTo := int32(repliedTo_32)
	repliesTo := int32(repliesTo_32)
	res, err := db.DeleteRepliesTo(ctx, repliesTo, repliedTo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"deleted_repliesTo": res,
	})
}
