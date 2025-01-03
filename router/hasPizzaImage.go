package router

import (
	"net/http"
	"server-api/db"

	"github.com/gin-gonic/gin"
)

func postHasPizzaImage(ctx *gin.Context) {
	var hasPizzaImage db.HasPizzaImage

	err := ctx.Bind(&hasPizzaImage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := db.CreateHasPizzaImage(hasPizzaImage, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"hasPizzaImage": res,
	})
}

func getHasPizzaImage(ctx *gin.Context) {
	name := ctx.Param("name")
	res, err := db.GetHasPizzaImage(name, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"hasPizzaImage": res,
	})
}

func deleteHasPizzaImage(ctx *gin.Context) {
	name := ctx.Param("name")
	res, err := db.DeleteHasPizzaImage(name, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"deleted_hasPizzaImage": res,
	})
}

func putHasPizzaImage(ctx *gin.Context) {
	var updatedHasPizzaImage db.HasPizzaImage

	err := ctx.Bind(&updatedHasPizzaImage)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	name := ctx.Param("name")

	res, err := db.UpdateHasPizzaImage(updatedHasPizzaImage, name, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"updated_hasPizzaImage": res,
	})
}
