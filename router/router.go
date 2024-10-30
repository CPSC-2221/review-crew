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
	r.GET("/user/:email", getUser)
	r.PUT("/user/:email", putUser)
	r.DELETE("/user/:email", deleteUser)
	r.POST("/create-db", db.CreateUsersTable)

	r.POST("/manages", postManages)
	r.GET("/manages/:email/:restaurantID", getManages)
	r.PUT("/manages/:email/:restaurantID", putManages)
	r.DELETE("/manages/:email/:restaurantID", deleteManages)
	r.POST("/create-managesTable", db.CreateManagesTable)

	return r
}
