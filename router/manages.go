package router

import (
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
