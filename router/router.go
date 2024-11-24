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
	r.GET("/ownerDashboard", ownerDashboard)

	r.POST("/createuser", createNewUser)
	r.POST("/loginuser", loginUser)

	r.POST("/addmanager", addManager)
	r.POST("/updatemanager", updateManager)
	r.POST("/deletemanager", deleteManager)

	r.POST("/updatedescription", putDescription)
	r.POST("/createRestaurant", createRestaurant)

	r.POST("/createreview", createReview)
	r.POST("/createreply", createReply)
	r.POST("/deletereview", removeReview)

	r.POST("/has-pizza-image", postHasPizzaImage)
	r.GET("/has-pizza-image/:name", getHasPizzaImage)
	r.PUT("/has-pizza-image/:name", putHasPizzaImage)
	r.DELETE("/has-pizza-image/:name", deleteHasPizzaImage)
	r.POST("/create-has-pizza-image-table", db.CreateHasPizzaImageTable)

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

func setup(ctx *gin.Context) {
	db.CreateUsersTable(ctx)
	db.CreateRestaurantTable(ctx)
	db.CreateHasPizzaImageTable(ctx)
	db.CreateHasBurgerEmojiTable(ctx)
	db.CreateReviewTable(ctx)
	db.CreateRepliesToTable(ctx)
	db.CreateManagesTable(ctx)
	db.CreateLikesTable(ctx)
	db.CreateOwnsTable(ctx)

	err := db.InsertDummyData(ctx)
	if err != nil {
		panic(err)
	}

	println("Finished setting up test rows")
}
