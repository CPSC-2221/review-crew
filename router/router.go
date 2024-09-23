package router

import (
	"server-api/db"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/user", postUser)
	r.GET("/user/:id", getUser)
	r.GET("/users", getUsers)
	r.PUT("/user/:id", putUser)
	r.DELETE("/user/:id", deleteUser)
	r.POST("/create-db", db.CreateUsersTable)

	return r
}
