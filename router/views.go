package router

import (
	"net/http"
	"server-api/db"
	"server-api/render"
	"server-api/views"

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

	render.Render(ctx, http.StatusOK, views.Index(res))
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

func getUsers(ctx *gin.Context) {
	res, err := db.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	render.Render(ctx, http.StatusOK, views.Users(res))
}
