package router

import (
	"server-api/db"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Static("static", "static")

	r.GET("/", index)
	r.GET("/users", renderUsers)
	r.GET("/location/:id", renderLocation)
	r.GET("/signup", signUp)
	r.GET("/login", logIn)
	r.GET("/logout", logout)
	r.POST("/openreply", openReply)
	r.GET("/createRestaurant", gotoCreateRestaurant)

	r.POST("/user", postUser)
	r.GET("/user/:email", getUser)
	r.PUT("/user/:email", putUser)
	r.DELETE("/user/:email", deleteUser)
	r.POST("/create-user-table", db.CreateUsersTable)
	r.POST("/createuser", createNewUser)
	r.POST("/loginuser", loginUser)

	r.POST("/owns", postOwn)
	r.GET("/owns/:email/:restaurantID", getOwn)
	r.PUT("/owns/:email/:restaurantID", putOwn)
	r.DELETE("/owns/:email/:restaurantID", deleteOwn)
	r.POST("/create-owns-table", db.CreateOwnsTable)

	r.POST("/manages", postManages)
	r.GET("/manages/:email/:restaurantID", getManages)
	r.PUT("/manages/:email/:restaurantID", putManages)
	r.DELETE("/manages/:email/:restaurantID", deleteManages)
	r.POST("/create-manages-table", db.CreateManagesTable)

	r.POST("/restaurant", postRestaurant)
	r.PUT("/restaurant/:id", putRestaurant)
	r.DELETE("/restaurant/:id", deleteRestaurant)
	r.POST("/create-restaurant-table", db.CreateRestaurantTable)
	r.POST("/updatedescription", putDescription)
	r.POST("/createRestaurant", createRestaurant)

	r.POST("/review", postReview)
	r.GET("/review/:id", getReview)
	r.PUT("/review/:id", putReview)
	r.DELETE("/review/:id", deleteReview)
	r.POST("/create-review-table", db.CreateReviewTable)
	r.POST("/createreview", createReview)
	r.POST("/createreply", createReply)
	r.POST("/deletereview", removeReview)

	r.POST("/replies-to", postRepliesTo)
	r.DELETE("/replies-to/:repliedTo/:repliesTo", deleteRepliesTo)
	r.POST("/create-replies-to-table", db.CreateRepliesToTable)

	r.POST("/has-pizza-image", postHasPizzaImage)
	r.GET("/has-pizza-image/:name", getHasPizzaImage)
	r.PUT("/has-pizza-image/:name", putHasPizzaImage)
	r.DELETE("/has-pizza-image/:name", deleteHasPizzaImage)
	r.POST("/create-has-pizza-image-table", db.CreateHasPizzaImageTable)

	r.POST("/like", postLike)
	r.GET("/like/:reviewID", getReviewLikes)
	r.DELETE("/like/:email/:reviewID", deleteLike)
	r.POST("/create-likes-table", db.CreateLikesTable)
	r.POST("/createlike", createLike)
	r.POST("/deletelike", removeLike)

	r.POST("/has-burger-emoji", postHasBurgerEmoji)
	r.GET("/has-burger-emoji/:username", getHasBurgerEmoji)
	r.PUT("/has-burger-emoji/:username", putHasBurgerEmoji)
	r.DELETE("/has-burger-emoji:username", deleteHasBurgerEmoji)
	r.POST("/create-has-burger-emoji-table", db.CreateHasBurgerEmojiTable)

	return r
}

func logout(ctx *gin.Context) {
	ctx.SetCookie("auth", "", 0, "", "", false, true)
	ctx.Header("HX-Refresh", "true")
}
