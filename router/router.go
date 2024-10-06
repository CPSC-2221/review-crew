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

	return r
}
