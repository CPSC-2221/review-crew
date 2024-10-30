package router

import (
	"net/http"
	"server-api/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func postOwn(ctx *gin.Context) {
	var own db.Own

	err := ctx.Bind(&own)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := db.CreateOwn(&own, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"own": res,
	})
}

func getOwn(ctx *gin.Context) {
	id_64, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := int32(id_64)

	email := ctx.Param("email")
	res, err := db.GetOwn(email, id, ctx)
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

func deleteOwn(ctx *gin.Context) {
	id_64, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := int32(id_64)

	email := ctx.Param("email")
	res, err := db.DeleteOwn(email, id, ctx)
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

func putOwn(ctx *gin.Context) {
	var updatedOwn db.Own

	err := ctx.Bind(&updatedOwn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id_64, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := int32(id_64)

	email := ctx.Param("email")
	dbown, err := db.GetOwn(email, id, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbown.Email = updatedOwn.Email
	dbown.RestaurantID = updatedOwn.RestaurantID

	res, err := db.UpdateOwn(dbown, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"updated_own": res,
	})
}
