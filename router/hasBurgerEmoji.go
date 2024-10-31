package router

import (
	"net/http"
	"server-api/db"

	"github.com/gin-gonic/gin"
)

func postHasBurgerEmoji(ctx *gin.Context) {
	var hasBurgerEmoji db.HasBurgerEmoji

	err := ctx.Bind(&hasBurgerEmoji)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := db.CreateHasBurgerEmoji(hasBurgerEmoji, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"burgerEmoji": res,
	})
}

func getHasBurgerEmoji(ctx *gin.Context) {
	username := ctx.Param("username")
	res, err := db.GetHasBurgerEmoji(username, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"burgerEmoji": res,
	})
}

func deleteHasBurgerEmoji(ctx *gin.Context) {
	username := ctx.Param("username")
	res, err := db.DeleteHasBurgerEmoji(username, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"deleted_burgerEmoji": res,
	})
}

func putHasBurgerEmoji(ctx *gin.Context) {
	var updatedHasBurgerEmoji db.HasBurgerEmoji

	err := ctx.Bind(&updatedHasBurgerEmoji)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	username := ctx.Param("username")

	res, err := db.UpdateHasBurgerEmoji(updatedHasBurgerEmoji, username, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"updated_burgerEmoji": res,
	})
}
