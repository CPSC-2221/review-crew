package router

import (
	"net/http"
	"server-api/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addManager(ctx *gin.Context) {
	var manages db.Manages

	r_id, _ := strconv.ParseInt(ctx.PostForm("restaurantID"), 10, 32)
	manages.RestaurantID = int32(r_id)

	manages.CanUpdateListing = ctx.PostForm("canUpdateListing") == "on"

	manages.CanDeleteComments = ctx.PostForm("canDeleteComments") == "on"

	manages.Email = db.GetEmailFromUsername(ctx.PostForm("username"), ctx)

	email, _ := ctx.Cookie("auth")
	if manages.Email == email {
		return
	}

	db.CreateManages(manages, ctx)
	ctx.Header("HX-Refresh", "true")
}

func deleteManager(ctx *gin.Context) {
	restaurantID, _ := strconv.ParseInt(ctx.PostForm("restaurantID"), 10, 32)
	email := ctx.PostForm("email")
	db.DeleteManages(email, int32(restaurantID), ctx)
	ctx.Header("HX-Refresh", "true")
}

func updateManager(ctx *gin.Context) {
	var manages db.Manages

	r_id, _ := strconv.ParseInt(ctx.PostForm("restaurantID"), 10, 32)
	manages.RestaurantID = int32(r_id)
	manages.CanUpdateListing = ctx.PostForm("canUpdateListing") == "on"
	manages.CanDeleteComments = ctx.PostForm("canDeleteComments") == "on"
	manages.Email = ctx.PostForm("email")

	println(manages.RestaurantID)
	println(manages.Email)

	db.UpdateManages(manages, ctx)
	ctx.Header("HX-Refresh", "true")
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
