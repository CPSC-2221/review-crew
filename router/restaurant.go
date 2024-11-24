package router

import (
	"net/http"
	"server-api/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	usr, _ := getUserFromCookie(ctx)
	println(usr.Email)
	resown, err := db.CreateOwn(db.Own{
		Email:        usr.Email,
		RestaurantID: res.ID,
	}, ctx)
	if err != nil {
		panic(err)
	}
	println(resown.Email)
	println(resown.RestaurantID)

	ctx.Header("HX-Redirect", "/location/"+strconv.Itoa(int(res.ID)))
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
