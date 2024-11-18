package router

import (
	"net/http"
	"server-api/db"
	"server-api/render"
	"server-api/views"
	"strconv"

	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	res, err := db.GetRestaurants(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	render.Render(ctx, http.StatusOK, views.Index(views.Home(res)))
}

func home(ctx *gin.Context) {
	res, err := db.GetRestaurants(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	render.Render(ctx, http.StatusOK, views.Home(res))
}

func renderUsers(ctx *gin.Context) {
	res, err := db.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if ctx.GetHeader("HX-Request") == "true" {
		render.Render(ctx, http.StatusOK, views.Users(res))
	} else {
		render.Render(ctx, http.StatusOK, views.Index(views.Users(res)))
	}
}

func renderLocation(ctx *gin.Context) {
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

	if ctx.GetHeader("HX-Request") == "true" {
		render.Render(ctx, http.StatusOK, views.Location(res))
	} else {
		render.Render(ctx, http.StatusOK, views.Index(views.Location(res)))
	}
}
