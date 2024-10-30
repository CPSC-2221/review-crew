package router

import (
	"net/http"
	"server-api/db"

	"github.com/gin-gonic/gin"
)

func postUser(ctx *gin.Context) {
	var user db.User

	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := db.CreateUser(&user, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"user": res,
	})
}

func getUser(ctx *gin.Context) {
	email := ctx.Param("email")
	res, err := db.GetUser(email, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": res,
	})
}

func deleteUser(ctx *gin.Context) {
	email := ctx.Param("email")
	res, err := db.DeleteUser(email, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"deleted_user": res,
	})
}

func putUser(ctx *gin.Context) {
	var updatedUser db.User

	err := ctx.Bind(&updatedUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	email := ctx.Param("email")

	dbuser, err := db.GetUser(email, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbuser.Username = updatedUser.Username
	dbuser.Email = updatedUser.Email

	res, err := db.UpdateUser(dbuser, ctx)
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
