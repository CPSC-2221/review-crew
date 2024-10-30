package router

import (
	"net/http"
	"server-api/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func postManages(ctx *gin.Context) {
	var manages db.Manages

	err := ctx.Bind(&manages)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := db.CreateManages(&manages, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"manages": res,
	})
}

func getManages(ctx *gin.Context) {
	email := ctx.Param("email")
	id_32, err := strconv.ParseInt(ctx.Param("restaurantID"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	restaurantID := int32(id_32)
	res, err := db.GetManages(email, restaurantID, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"manages": res,
	})
}

func deleteManages(ctx *gin.Context) {
	email := ctx.Param("email")
	id_32, err := strconv.ParseInt(ctx.Param("restaurantID"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	restaurantID := int32(id_32)
	res, err := db.DeleteManages(email, restaurantID, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"deleted_manages": res,
	})
}

func putManages(ctx *gin.Context) {
	var updatedManages db.Manages

	err := ctx.Bind(&updatedManages)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	email := ctx.Param("email")
	id_32, err := strconv.ParseInt(ctx.Param("restaurantID"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	restaurantID := int32(id_32)

	dbmanages, err := db.GetManages(email, restaurantID, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbmanages.Email = updatedManages.Email
	dbmanages.RestaurantID = updatedManages.RestaurantID
	dbmanages.CanDeleteComments = updatedManages.CanDeleteComments
	dbmanages.CanUpdateListing = updatedManages.CanUpdateListing

	res, err := db.UpdateManages(dbmanages, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"updated_user": res,
	})

}
