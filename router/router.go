package router

import (
	"server-api/db"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Static("static", "static")

	r.GET("/", index)
	r.GET("/users", getUsers)
	r.GET("/home", home)

	r.POST("/user", postUser)
	r.GET("/user/:id", getUser)
	r.PUT("/user/:id", putUser)
	r.DELETE("/user/:id", deleteUser)
	r.POST("/create-db", db.CreateUsersTable)

	r.POST("/owns", postOwn)
	r.GET("/owns/:email/:restaurantID", getOwn)
	r.PUT("/owns/:email/:restaurantID", putOwn)
	r.DELETE("/owns/:email/:restaurantID", deleteOwn)
	r.POST("/create-ownsTable", db.CreateOwnsTable)

	r.POST("/create-restaurantTable", db.CreateRestaurantTable)
	return r
}
