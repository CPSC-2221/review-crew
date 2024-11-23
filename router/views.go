package router

import (
	"net/http"
	"server-api/db"
	"server-api/render"
	"server-api/views"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserFromCookie(ctx *gin.Context) (*db.User, error) {
	var usr *db.User = nil

	email, err := ctx.Cookie("auth")
	if err != nil {
		println("error fetching cookie: " + err.Error())
		return nil, err
	}

	usr, err = db.GetUser(email, ctx)
	if err != nil {
		println("error fetching user: " + err.Error())
		return nil, err
	}

	return usr, nil
}

func index(ctx *gin.Context) {
	usr, _ := getUserFromCookie(ctx)
	renderIndex(ctx, usr)
}

func renderIndex(ctx *gin.Context, usr *db.User) {
	res, err := db.GetRestaurants(ctx)
	if err != nil {
		setup(ctx)
	}

	render.Render(ctx, http.StatusOK, views.Index(views.Home(res), usr))
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
	res, _ := db.GetUsers(ctx)

	if ctx.GetHeader("HX-Request") == "true" {
		render.Render(ctx, http.StatusOK, views.Users(res))
	} else {
		usr, _ := getUserFromCookie(ctx)
		render.Render(ctx, http.StatusOK, views.Index(views.Users(res), usr))
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

	usr, _ := getUserFromCookie(ctx)
	if ctx.GetHeader("HX-Request") == "true" {
		render.Render(ctx, http.StatusOK, views.Location(res, usr))
	} else {
		render.Render(ctx, http.StatusOK, views.Index(views.Location(res, usr), usr))
	}
}

func signUp(ctx *gin.Context) {
	usr, _ := getUserFromCookie(ctx)
	if usr != nil {
		renderIndex(ctx, usr)
		return
	}

	if ctx.GetHeader("HX-Request") == "true" {
		render.Render(ctx, http.StatusOK, views.SignUp())
	} else {
		render.Render(ctx, http.StatusOK, views.Index(views.SignUp(), usr))
	}
}

func renderSignUp(ctx *gin.Context, errors ...string) {
	render.Render(ctx, http.StatusOK, views.SignUp(errors...))
}

func logIn(ctx *gin.Context) {
	usr, _ := getUserFromCookie(ctx)
	if usr != nil {
		renderIndex(ctx, usr)
		return
	}

	if ctx.GetHeader("HX-Request") == "true" {
		render.Render(ctx, http.StatusOK, views.LogIn())
	} else {
		render.Render(ctx, http.StatusOK, views.Index(views.LogIn(), usr))
	}
}

func renderLogIn(ctx *gin.Context, errors ...string) {
	render.Render(ctx, http.StatusOK, views.LogIn(errors...))
}

func openReply(ctx *gin.Context) {
	usr, _ := getUserFromCookie(ctx)
	if usr == nil {
		return
	}
	render.Render(ctx, http.StatusOK, views.ReplyInput(usr.Email, ctx.PostForm("locationID"), ctx.PostForm("reviewID")))
}

func gotoCreateRestaurant(ctx *gin.Context) {
	usr, _ := getUserFromCookie(ctx)
	if usr == nil {
		renderIndex(ctx, usr)
		return
	}

	if ctx.GetHeader("HX-Request") == "true" {
		render.Render(ctx, http.StatusOK, views.CreateRestaurant())
	} else {
		render.Render(ctx, http.StatusOK, views.Index(views.CreateRestaurant(), usr))
	}
}

func renderCreateRestaurant(ctx *gin.Context, errors ...string) {
	render.Render(ctx, http.StatusOK, views.CreateRestaurant(errors...))
}

func ownerDashboard(ctx *gin.Context) {
	usr, _ := getUserFromCookie(ctx)
	if usr == nil {
		renderIndex(ctx, usr)
		return
	}

	if ctx.GetHeader("HX-Request") == "true" {
		render.Render(ctx, http.StatusOK, views.OwnerDashboard(usr))
	} else {
		render.Render(ctx, http.StatusOK, views.Index(views.OwnerDashboard(usr), usr))
	}
}
