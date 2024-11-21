package router

import (
	"net/http"
	"server-api/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func createRestaurant(ctx *gin.Context) {
	var restaurant db.Restaurant
	restaurant.Name = ctx.PostForm("name")
	restaurant.Location = ctx.PostForm("location")
	restaurant.Description = ctx.PostForm("description")
	errors := make([]string, 0)

	if len(restaurant.Name) < 1 {
		errors = append(errors, "ERROR: Restaurant name must have at least 1 character")
	}

	if len(restaurant.Location) < 1 {
		errors = append(errors, "ERROR: Restaurant location must have at least 1 character")
	}

	if len(restaurant.Description) < 10 {
		errors = append(errors, "ERROR: Restaurant description must have at least 10 character")
	}

	if len(errors) > 0 {
		renderCreateRestaurant(ctx, errors...)
		return
	}

	res, err := db.CreateRestaurant(&restaurant, ctx)
	if err != nil {
		renderCreateRestaurant(ctx, err.Error())
		return
	}

	ctx.Header("HX-Redirect", "/location/"+strconv.Itoa(int(res.ID)))
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

func putDescription(ctx *gin.Context) {
	desc := ctx.PostForm("description")

	id_32, err := strconv.ParseInt(ctx.PostForm("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := int32(id_32)

	db.UpdateDescription(desc, id, ctx)

	ctx.Header("HX-Refresh", "true")
}
