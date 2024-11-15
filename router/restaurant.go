package router

import (
	"net/http"
	"server-api/db"
	"server-api/render"
	"server-api/views"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getLocation(ctx *gin.Context) {
	id_32, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := int32(id_32)

	res, err := db.GetRestaurant(id, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	render.Render(ctx, http.StatusOK, views.Location(res))
}

func postRestaurant(ctx *gin.Context) {
	var restaurant db.Restaurant

	err := ctx.Bind(&restaurant)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := db.CreateRestaurant(&restaurant, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"restaurant": res,
	})
}

func deleteRestaurant(ctx *gin.Context) {
	id_32, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := int32(id_32)

	res, err := db.DeleteRestaurant(id, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"deleted_restaurant": res,
	})
}

func putRestaurant(ctx *gin.Context) {
	var updated_restaurant db.Restaurant

	err := ctx.Bind(&updated_restaurant)
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

	updated_restaurant.ID = id

	res, err := db.UpdateRestaurant(&updated_restaurant, id, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"updated_restaurant": res,
	})
}
