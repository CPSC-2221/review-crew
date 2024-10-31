package router

import (
	"net/http"
	"server-api/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func postLimit(ctx *gin.Context) {
	var limit db.ReviewCharacterLimit

	err := ctx.Bind(&limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateReviewCharacterLimit(limit, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"limit": res,
	})
}

func getLimit(ctx *gin.Context) {
	restaurantID_32, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	restaurantID := int32(restaurantID_32)

	res, err := db.GetReviewCharacterLimit(restaurantID, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"limit": res,
	})
}

func deleteLimit(ctx *gin.Context) {
	restaurantID_32, err := strconv.ParseInt(ctx.Param("restaurantID"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid restaurant ID",
		})
		return
	}
	restaurantID := int32(restaurantID_32)

	res, err := db.DeleteReviewCharacterLimit(restaurantID, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"deleted_limit": res,
	})
}

func putLimit(ctx *gin.Context) {
	var updatedLimit db.ReviewCharacterLimit

	err := ctx.Bind(&updatedLimit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	restaurantID_32, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	restaurantID := int32(restaurantID_32)
	updatedLimit.RestaurantID = restaurantID

	res, err := db.UpdateReviewCharacterLimit(updatedLimit, restaurantID, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"updated_limit": res,
	})
}
