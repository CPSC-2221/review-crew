package router

import (
	"server-api/db"

	"github.com/gin-gonic/gin"
)

func createNewUser(ctx *gin.Context) {
	var usr db.User
	usr.Email = ctx.PostForm("email")
	usr.Username = ctx.PostForm("username")
	errors := make([]string, 0)

	if len(usr.Email) < 5 {
		errors = append(errors, "ERROR: email must be 5 characters or longer")
	}

	if len(usr.Username) < 4 {
		errors = append(errors, "ERROR: username must be 4 characters or longer")
	}

	if len(errors) > 0 {
		renderSignUp(ctx, errors...)
		return
	}

	_, err := db.CreateUser(usr, ctx)
	if err != nil {
		renderSignUp(ctx, err.Error())
		return
	}

	ctx.SetCookie("auth", usr.Email, 2592000, "", "", false, true)
	ctx.Header("HX-Redirect", "/")
}

func loginUser(ctx *gin.Context) {
	loginEmail := ctx.PostForm("email")
	loginUsername := ctx.PostForm("username")

	usr, _ := db.GetUser(loginEmail, ctx)
	if usr == nil {
		renderLogIn(ctx, "User account does not exist")
		return
	}

	if usr.Username != loginUsername {
		renderLogIn(ctx, "Username is incorrect")
		return
	}

	ctx.SetCookie("auth", usr.Email, 2592000, "", "", false, true)
	ctx.Header("HX-Redirect", "/")
}

func logout(ctx *gin.Context) {
	ctx.SetCookie("auth", "", 0, "", "", false, true)
	ctx.Header("HX-Refresh", "true")
}
